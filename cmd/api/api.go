package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Replais/replais-api/internal/service"
)

type config struct {
	addr     string
	dbConfig dbConfig
}

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  time.Duration
}

type application struct {
	config
	services service.Services
}

func (a *application) run(h http.Handler) error {
	srv := http.Server{
		Addr:         a.config.addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}
	log.Printf("Server has started at %s", a.config.addr)
	return srv.ListenAndServe()
}
