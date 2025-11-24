package service

import (
	"context"

	"github.com/Replais/replais-api/internal/logger"
	"github.com/Replais/replais-api/internal/model"
	"github.com/Replais/replais-api/internal/store"
)

// ContactsService defines business logic for contacts.
type ContactsService interface {
	Create(ctx context.Context, contact *model.Contact) error
	GetContactSettings(ctx context.Context, userID, platform, contactKey string) (*model.ContactSetting, *model.Persona, error)
	// later: List, Update, etc.
}

type contactsService struct {
	contacts store.Contacts
	personas store.Personas
	logger   logger.Logger
}

func NewContactsService(contacts store.Contacts, personas store.Personas, log logger.Logger) ContactsService {
	return &contactsService{
		contacts: contacts,
		personas: personas,
		logger:   log,
	}
}

func (s *contactsService) Create(ctx context.Context, contact *model.Contact) error {
	s.logger.Info("Creating contact: %s for user %s", contact.DisplayName, contact.UserID)
	// place for validation, dedupe, etc.
	if err := s.contacts.Create(ctx, contact); err != nil {
		s.logger.Error("Failed to create contact %s: %v", contact.DisplayName, err)
		return err
	}
	s.logger.Info("Successfully created contact: %s", contact.DisplayName)
	return nil
}

// GetContactSettings orchestrates fetching contact by key, then contact settings, then persona
// This is the business logic layer - it combines data from multiple stores
// If contact settings are not found, returns default settings with empty instructions and default persona
func (s *contactsService) GetContactSettings(ctx context.Context, userID, platform, contactKey string) (*model.ContactSetting, *model.Persona, error) {
	s.logger.Info("Getting contact settings for user %s, platform %s, contactKey %s", userID, platform, contactKey)

	// Step 1: Get contact by key
	contact, err := s.contacts.GetByKey(ctx, userID, platform, contactKey)
	if err != nil {
		s.logger.Error("Failed to get contact: %v", err)
		return nil, nil, err
	}

	// Step 2: Get contact settings by contact ID
	contactSettings, err := s.contacts.GetContactSettingsByContactID(ctx, contact.ID)
	if err != nil {
		// If contact settings not found, return default settings
		if err == store.ErrNotFound {
			s.logger.Info("Contact settings not found for contact %s, returning default settings", contact.ID)

			// Get default persona
			defaultPersona, err := s.personas.GetDefault(ctx)
			if err != nil {
				s.logger.Error("Failed to get default persona: %v", err)
				return nil, nil, err
			}

			// Create default contact settings with empty instructions
			defaultSettings := &model.ContactSetting{
				ID:           "", // No ID since it doesn't exist in DB
				UserID:       contact.UserID,
				ContactID:    contact.ID,
				PersonaID:    defaultPersona.ID,
				Instructions: "",
				Extra:        make(map[string]any),
			}

			return defaultSettings, defaultPersona, nil
		}
		s.logger.Error("Failed to get contact settings: %v", err)
		return nil, nil, err
	}

	// Step 3: Get persona by ID (if persona_id exists)
	var persona *model.Persona
	if contactSettings.PersonaID != "" {
		persona, err = s.personas.GetByID(ctx, contactSettings.PersonaID)
		if err != nil {
			s.logger.Error("Failed to get persona: %v", err)
			return nil, nil, err
		}
	} else {
		// If no persona_id is set, use default persona
		defaultPersona, err := s.personas.GetDefault(ctx)
		if err != nil {
			s.logger.Error("Failed to get default persona: %v", err)
			return nil, nil, err
		}
		persona = defaultPersona
	}

	s.logger.Info("Successfully retrieved contact settings for contact %s", contact.ID)
	return contactSettings, persona, nil
}
