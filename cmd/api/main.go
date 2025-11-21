package main

import (
	"log"

	"github.com/Replais/replais-api/internal/env"
	"github.com/Replais/replais-api/internal/service"
	"github.com/Replais/replais-api/internal/store/postgres"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	// TODO: open real DB here
	store := postgres.NewStorage(nil)
	services := service.NewServices(store)

	app := &application{
		config:   cfg,
		services: services,
	}

	if err := app.run(app.routes()); err != nil {
		log.Fatal(err)
	}
}
