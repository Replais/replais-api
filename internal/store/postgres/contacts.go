package postgres

import (
	"context"
	"database/sql"
)

// ContactsStore implements the store.Contacts interface for Postgres
type ContactsStore struct {
	db *sql.DB
}

// NewContactsStore creates a new ContactsStore instance
func NewContactsStore(db *sql.DB) *ContactsStore {
	return &ContactsStore{db: db}
}

// Create implements the store.Contacts interface
func (s *ContactsStore) Create(ctx context.Context) error {
	// TODO: Implement actual database logic
	return nil
}
