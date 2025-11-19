package postgres

import (
	"database/sql"

	"github.com/Replais/replais-api/internal/store"
)

// NewStorage creates a new Postgres storage implementation
func NewStorage(db *sql.DB) store.Storage {
	return store.Storage{
		Users:    NewUsersStore(db),
		Contacts: NewContactsStore(db),
	}
}
