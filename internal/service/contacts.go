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
	// later: List, Update, etc.
}

type contactsService struct {
	contacts store.Contacts
	logger   logger.Logger
}

func NewContactsService(contacts store.Contacts, log logger.Logger) ContactsService {
	return &contactsService{
		contacts: contacts,
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
