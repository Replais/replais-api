package postgres

import (
	"database/sql"

	"github.com/Replais/replais-api/internal/logger"
	"github.com/Replais/replais-api/internal/store"
)

// Storage implements the store.Storage interface for Postgres
type PostgresStorage struct {
	users    *UsersStore
	contacts *ContactsStore
	personas *PersonasStore
}

// NewStorage creates a new Postgres storage implementation
// Logger is passed here so stores can log database operations
func NewStorage(db *sql.DB, log logger.Logger) store.Storage {
	return &PostgresStorage{
		users:    NewUsersStore(db, log),
		contacts: NewContactsStore(db, log),
		personas: NewPersonasStore(db, log),
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

// Personas returns the personas repository
func (s *PostgresStorage) Personas() store.Personas {
	return s.personas
}
