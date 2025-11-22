package main

import (
	"log"
	"time"

	"github.com/Replais/replais-api/internal/db"
	"github.com/Replais/replais-api/internal/env"
	"github.com/Replais/replais-api/internal/service"
	"github.com/Replais/replais-api/internal/store/postgres"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		dbConfig: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://replais_user:replais_password@localhost:5432/replais?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetDuration("DB_MAX_IDLE_TIME", 15*time.Minute),
		},
	}

	db, err := db.New(cfg.dbConfig.addr, cfg.dbConfig.maxOpenConns, cfg.dbConfig.maxIdleConns, cfg.dbConfig.maxIdleTime)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	log.Println("DB connection successful")
	store := postgres.NewStorage(db)
	services := service.NewServices(store)

	app := &application{
		config:   cfg,
		services: services,
	}

	if err := app.run(app.routes()); err != nil {
		log.Fatal(err)
	}
}
