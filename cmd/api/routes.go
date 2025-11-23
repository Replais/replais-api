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
	// Controllers get logger (for request logging) and config (for server-specific needs)
	health := controllers.NewHealthController(app.logger, app.config)
	contacts := controllers.NewContactsController(app.services.Contacts, app.logger, app.config)
	users := controllers.NewUserController(app.services.Users, app.logger, app.config)
	personas := controllers.NewPersonasController(app.services.Personas, app.logger)
	// health
	r.Get("/v1/health", health.Health)
	r.Route("/v1", func(r chi.Router) {
		// users
		r.Route("/users", func(r chi.Router) {
			r.Post("/", users.Create)
			// r.Get("/{id}", app.getUserHandler)
		})
		// contacts
		r.Route("/contacts", func(r chi.Router) {
			r.Post("/", contacts.Create)
		})
		// personas
		r.Route("/personas", func(r chi.Router) {
			r.Post("/", personas.Create)
			r.Get("/", personas.GetAll) // Public route: returns only slug and label
		})
		// admin/internal routes
		r.Route("/admin", func(r chi.Router) {
			r.Route("/personas", func(r chi.Router) {
				r.Get("/", personas.GetAllAdmin) // Admin route: returns full persona details
			})
		})
	})

	return r
}
