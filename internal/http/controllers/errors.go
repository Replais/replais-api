package controllers

import (
	"net/http"

	"github.com/Replais/replais-api/internal/logger"
	"github.com/go-chi/chi/v5/middleware"
)

func internalServerError(logger logger.Logger, w http.ResponseWriter, r *http.Request, err error) {
	logger.Error("Internal server error: %v", err, map[string]any{
		"path":     r.URL.Path,
		"method":   r.Method,
		"status":   http.StatusInternalServerError,
		"trace_id": middleware.GetReqID(r.Context()),
	})
	writeJSONError(w, http.StatusInternalServerError, "The server encountered a problem and could not process your request")
}

func badRequestError(logger logger.Logger, w http.ResponseWriter, r *http.Request, err error) {
	logger.Error("Bad request: %v", err, map[string]any{
		"path":     r.URL.Path,
		"method":   r.Method,
		"status":   http.StatusBadRequest,
		"trace_id": middleware.GetReqID(r.Context()),
	})
	writeJSONError(w, http.StatusBadRequest, "The request is invalid")
}

func notFoundError(logger logger.Logger, w http.ResponseWriter, r *http.Request, err error) {
	logger.Error("Not found: %v", err, map[string]any{
		"path":     r.URL.Path,
		"method":   r.Method,
		"status":   http.StatusNotFound,
		"trace_id": middleware.GetReqID(r.Context()),
	})
	writeJSONError(w, http.StatusNotFound, "The requested resource was not found")
}
