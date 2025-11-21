package service

import (
	"context"

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
}

func NewContactsService(contacts store.Contacts) ContactsService {
	return &contactsService{contacts: contacts}
}

func (s *contactsService) Create(ctx context.Context, contact *model.Contact) error {
	// place for validation, dedupe, etc.
	return s.contacts.Create(ctx, contact)
}
