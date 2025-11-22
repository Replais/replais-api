package main

import (
	"net/http"
	"time"

	"github.com/Replais/replais-api/internal/config"
	"github.com/Replais/replais-api/internal/logger"
	"github.com/Replais/replais-api/internal/service"
)

type application struct {
	config   config.Config
	logger   logger.Logger
	services service.Services
}

func (a *application) run(h http.Handler) error {
	srv := http.Server{
		Addr:         a.config.Addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}
	a.logger.Info("Server has started at %s", a.config.Addr)
	return srv.ListenAndServe()
}
