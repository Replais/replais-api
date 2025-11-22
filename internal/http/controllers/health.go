package controllers

import (
	"net/http"

	"github.com/Replais/replais-api/internal/config"
	"github.com/Replais/replais-api/internal/logger"
)

type HealthController struct {
	logger logger.Logger
	config config.Config
}

func NewHealthController(log logger.Logger, cfg config.Config) *HealthController {
	return &HealthController{
		logger: log,
		config: cfg,
	}
}

func (c *HealthController) Health(w http.ResponseWriter, r *http.Request) {
	c.logger.Debug("Health check requested")
	data := map[string]string{
		"status": "ok",
	}
	if err := jsonResponse(w, http.StatusOK, data); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
	}
}
