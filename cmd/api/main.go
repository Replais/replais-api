package main

import (
	"time"

	"github.com/Replais/replais-api/internal/config"
	"github.com/Replais/replais-api/internal/db"
	"github.com/Replais/replais-api/internal/env"
	"github.com/Replais/replais-api/internal/logger"
	"github.com/Replais/replais-api/internal/service"
	"github.com/Replais/replais-api/internal/store/postgres"
)

func main() {
	// Initialize logger first (so we can use it for other initialization)
	log := logger.NewStdLogger()

	// Load configuration
	cfg := config.Config{
		Addr: env.GetString("ADDR", ":8080"),
		DB: config.DBConfig{
			Addr:         env.GetString("DB_ADDR", "postgres://replais_user:replais_password@localhost:5432/replais?sslmode=disable"),
			MaxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			MaxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			MaxIdleTime:  env.GetDuration("DB_MAX_IDLE_TIME", 15*time.Minute),
		},
	}

	// Initialize database
	database, err := db.New(cfg.DB.Addr, cfg.DB.MaxOpenConns, cfg.DB.MaxIdleConns, cfg.DB.MaxIdleTime)
	if err != nil {
		log.Fatal("Failed to connect to database: %v", err)
	}
	defer database.Close()
	log.Info("DB connection successful")

	// Initialize store
	store := postgres.NewStorage(database, log)

	// Initialize services with logger
	services := service.NewServices(store, log)

	// Initialize application
	app := &application{
		config:   cfg,
		logger:   log,
		services: services,
	}

	// Start server
	if err := app.run(app.routes()); err != nil {
		log.Fatal("Server failed: %v", err)
	}
}
