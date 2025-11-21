package postgres

import (
	"database/sql"

	"github.com/Replais/replais-api/internal/store"
)

// Storage implements the store.Storage interface for Postgres
type PostgresStorage struct {
	users    *UsersStore
	contacts *ContactsStore
}

// NewStorage creates a new Postgres storage implementation
func NewStorage(db *sql.DB) store.Storage {
	return &PostgresStorage{
		users:    NewUsersStore(db),
		contacts: NewContactsStore(db),
	}
}

// Users returns the users repository
func (s *PostgresStorage) Users() store.Users {
	return s.users
}

// Contacts returns the contacts repository
func (s *PostgresStorage) Contacts() store.Contacts {
	return s.contacts
}
