// internal/http/controllers/contacts.go
package controllers

import (
	"net/http"

	"github.com/Replais/replais-api/internal/config"
	"github.com/Replais/replais-api/internal/logger"
	"github.com/Replais/replais-api/internal/model"
	"github.com/Replais/replais-api/internal/service"
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
