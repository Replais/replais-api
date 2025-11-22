package postgres

import (
	"context"
	"database/sql"

	"github.com/Replais/replais-api/internal/logger"
	"github.com/Replais/replais-api/internal/model"
)

// ContactsStore implements the store.Contacts interface for Postgres
type ContactsStore struct {
	db     *sql.DB
	logger logger.Logger
}

// NewContactsStore creates a new ContactsStore instance
func NewContactsStore(db *sql.DB, log logger.Logger) *ContactsStore {
	return &ContactsStore{
		db:     db,
		logger: log,
	}
}

// Create implements the store.Contacts interface
func (s *ContactsStore) Create(ctx context.Context, contact *model.Contact) error {
	query := `
		INSERT INTO contacts (id, user_id, platform, platform_contact_key, display_name, is_group)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at, updated_at
	`

	err := s.db.QueryRowContext(ctx, query, contact.ID, contact.UserID, contact.Platform, contact.PlatformContactKey, contact.DisplayName, contact.IsGroup).Scan(&contact.ID, &contact.CreatedAt, &contact.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}
