package postgres

import (
	"context"
	"database/sql"
)

// UsersStore implements the store.Users interface for Postgres
type UsersStore struct {
	db *sql.DB
}

// NewUsersStore creates a new UsersStore instance
func NewUsersStore(db *sql.DB) *UsersStore {
	return &UsersStore{db: db}
}

// Create implements the store.Users interface
func (s *UsersStore) Create(ctx context.Context) error {
	// TODO: Implement actual database logic
	return nil
}
