package main

import (
	"net/http"
	"time"

	"github.com/Replais/replais-api/internal/http/controllers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	r := chi.NewRouter()

	// global middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	// Instantiate controllers with DI
	health := controllers.NewHealthController()
	contacts := controllers.NewContactsController(app.services.Contacts)
	users := controllers.NewUserController(app.services.Users)

	// health
	r.Get("/v1/health", health.Health)
	r.Route("/v1", func(r chi.Router) {
		// users
		r.Route("/users", func(r chi.Router) {
			r.Post("/", users.Create)
			// r.Get("/{id}", app.getUserHandler)
			// r.Put("/{id}", app.updateUserHandler)
		})
		// contacts
		r.Route("/contacts", func(r chi.Router) {
			r.Post("/", contacts.Create)
			// r.Get("/", app.listContactsHandler)
			// r.Get("/{id}", app.getContactHandler)
		})
	})

	return r
}
