package postgres

import (
	"context"
	"database/sql"

	"github.com/Replais/replais-api/internal/logger"
	"github.com/Replais/replais-api/internal/model"
	"github.com/Replais/replais-api/internal/store"
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
		INSERT INTO contacts (user_id, platform, platform_contact_key, display_name, is_group)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, updated_at
	`

	err := s.db.QueryRowContext(ctx, query, contact.UserID, contact.Platform, contact.PlatformContactKey, contact.DisplayName, contact.IsGroup).Scan(&contact.ID, &contact.CreatedAt, &contact.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

// GetByKey retrieves a contact by user_id, platform, and platform_contact_key
func (s *ContactsStore) GetByKey(ctx context.Context, userID, platform, contactKey string) (*model.Contact, error) {
	query := `
		SELECT id, user_id, platform, platform_contact_key, display_name, is_group, created_at, updated_at
		FROM contacts
		WHERE user_id = $1 AND platform = $2 AND platform_contact_key = $3
	`
	contact := &model.Contact{}
	err := s.db.QueryRowContext(ctx, query, userID, platform, contactKey).Scan(
		&contact.ID, &contact.UserID, &contact.Platform, &contact.PlatformContactKey,
		&contact.DisplayName, &contact.IsGroup, &contact.CreatedAt, &contact.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrNotFound
		}
		return nil, err
	}
	return contact, nil
}

// GetContactSettingsByContactID retrieves contact settings by contact_id
func (s *ContactsStore) GetContactSettingsByContactID(ctx context.Context, contactID string) (*model.ContactSetting, error) {
	query := `
		SELECT id, user_id, contact_id, persona_id, instructions, extra, created_at, updated_at
		FROM contact_settings
		WHERE contact_id = $1
	`
	settings := &model.ContactSetting{}
	var personaID sql.NullString
	err := s.db.QueryRowContext(ctx, query, contactID).Scan(
		&settings.ID, &settings.UserID, &settings.ContactID, &personaID,
		&settings.Instructions, &settings.Extra, &settings.CreatedAt, &settings.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrNotFound
		}
		return nil, err
	}
	if personaID.Valid {
		settings.PersonaID = personaID.String
	}
	return settings, nil
}
