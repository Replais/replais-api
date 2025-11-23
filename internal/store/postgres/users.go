package postgres

import (
	"context"
	"database/sql"

	"github.com/Replais/replais-api/internal/logger"
	"github.com/Replais/replais-api/internal/model"
)

// UsersStore implements the store.Users interface for Postgres
type UsersStore struct {
	db     *sql.DB
	logger logger.Logger
}

// NewUsersStore creates a new UsersStore instance
func NewUsersStore(db *sql.DB, log logger.Logger) *UsersStore {
	return &UsersStore{
		db:     db,
		logger: log,
	}
}

// Create implements the store.Users interface
func (s *UsersStore) Create(ctx context.Context, user *model.User) error {
	query := `
		INSERT INTO users (email, name)
		VALUES ($1, $2)
		RETURNING id, created_at, updated_at
	`
	err := s.db.QueryRowContext(ctx, query, user.Email, user.Name).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}
