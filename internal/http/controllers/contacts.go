// internal/http/controllers/contacts.go
package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Replais/replais-api/internal/model"
	"github.com/Replais/replais-api/internal/service"
)

type ContactsController struct {
	contacts service.ContactsService
}

func NewContactsController(s service.ContactsService) *ContactsController {
	return &ContactsController{
		contacts: s,
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

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
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
		http.Error(w, "failed to create contact", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(contact)
}
