// internal/http/controllers/contacts.go
package controllers

import (
	"errors"
	"net/http"

	"github.com/Replais/replais-api/internal/config"
	"github.com/Replais/replais-api/internal/http/dto"
	"github.com/Replais/replais-api/internal/logger"
	"github.com/Replais/replais-api/internal/model"
	"github.com/Replais/replais-api/internal/service"
	"github.com/Replais/replais-api/internal/store"
)

type ContactsController struct {
	contacts service.ContactsService
	logger   logger.Logger
	config   config.Config
}

func NewContactsController(s service.ContactsService, log logger.Logger, cfg config.Config) *ContactsController {
	return &ContactsController{
		contacts: s,
		logger:   log,
		config:   cfg,
	}
}

// DTO for incoming payload
type createContactRequest struct {
	UserID             string `json:"user_id"`
	Platform           string `json:"platform"`
	PlatformContactKey string `json:"platform_contact_key"`
	DisplayName        string `json:"display_name"`
	IsGroup            bool   `json:"is_group"`
}

func (c *ContactsController) Create(w http.ResponseWriter, r *http.Request) {
	var req createContactRequest

	if err := readJSON(w, r, &req); err != nil {
		badRequestError(c.logger, w, r, err)
		return
	}

	contact := &model.Contact{
		UserID:             req.UserID,
		Platform:           req.Platform,
		PlatformContactKey: req.PlatformContactKey,
		DisplayName:        req.DisplayName,
		IsGroup:            req.IsGroup,
	}

	if err := c.contacts.Create(r.Context(), contact); err != nil {
		internalServerError(c.logger, w, r, err)
		return
	}

	err := jsonResponse(w, http.StatusCreated, contact)
	if err != nil {
		internalServerError(c.logger, w, r, err)
		return
	}

}

func (c *ContactsController) GetContactSettings(w http.ResponseWriter, r *http.Request) {
	// Extract query parameters
	userID := r.URL.Query().Get("user_id")
	platform := r.URL.Query().Get("platform")
	contactKey := r.URL.Query().Get("contact_key")

	// Validate required parameters
	if userID == "" || platform == "" || contactKey == "" {
		badRequestError(c.logger, w, r, errors.New("missing required parameters: user_id, platform, and contact_key are required"))
		return
	}

	// Call service to get contact settings and persona
	contactSettings, persona, err := c.contacts.GetContactSettings(r.Context(), userID, platform, contactKey)
	if err != nil {
		if err == store.ErrNotFound {
			notFoundError(c.logger, w, r, err)
			return
		}
		internalServerError(c.logger, w, r, err)
		return
	}

	// Service guarantees persona will always be returned (either from settings or default)
	// Map to DTO using the mapper
	response := dto.ToContactSettingsResponse(*contactSettings, *persona)
	err = jsonResponse(w, http.StatusOK, response)
	if err != nil {
		internalServerError(c.logger, w, r, err)
		return
	}
}
