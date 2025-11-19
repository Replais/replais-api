package main

import (
	"github.com/Replais/replais-api/internal/env"
	"log"
)

func main() {

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}
	app := &application{
		config: cfg,
	}
	err := app.run(app.mount())
	if err != nil {
		log.Fatal(err.Error())
	}

}
