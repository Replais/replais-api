package main

import (
	"log"

	"github.com/Replais/replais-api/internal/env"
	"github.com/Replais/replais-api/internal/store/postgres"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}
	store := postgres.NewStorage(nil)
	app := &application{
		config: cfg,
		store:  store,
	}

	err := app.run(app.mount())
	if err != nil {
		log.Fatal(err.Error())
	}

}
