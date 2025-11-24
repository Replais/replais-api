package postgres

import (
	"context"
	"database/sql"

	"github.com/Replais/replais-api/internal/logger"
	"github.com/Replais/replais-api/internal/model"
	"github.com/Replais/replais-api/internal/store"
)

type PersonasStore struct {
	db     *sql.DB
	logger logger.Logger
}

func NewPersonasStore(db *sql.DB, log logger.Logger) *PersonasStore {
	return &PersonasStore{
		db:     db,
		logger: log,
	}
}

func (s *PersonasStore) Create(ctx context.Context, persona *model.Persona) error {
	query := `
		INSERT INTO personas (user_id, slug, label, description, prompt_template, is_default)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at, updated_at
	`
	err := s.db.QueryRowContext(ctx, query, persona.UserID, persona.Slug, persona.Label, persona.Description, persona.PromptTemplate, persona.IsDefault).Scan(&persona.ID, &persona.CreatedAt, &persona.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (s *PersonasStore) GetAll(ctx context.Context) ([]model.Persona, error) {
	query := `
		SELECT id, user_id, slug, label, description, prompt_template, is_default, created_at, updated_at
		FROM personas
		WHERE user_id IS NULL
	`
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	personas := []model.Persona{}
	for rows.Next() {
		var persona model.Persona
		err := rows.Scan(&persona.ID, &persona.UserID, &persona.Slug, &persona.Label, &persona.Description, &persona.PromptTemplate, &persona.IsDefault, &persona.CreatedAt, &persona.UpdatedAt)
		if err != nil {
			return nil, err
		}
		personas = append(personas, persona)
	}
	return personas, nil
}

// GetByID retrieves a persona by ID
func (s *PersonasStore) GetByID(ctx context.Context, id string) (*model.Persona, error) {
	query := `
		SELECT id, user_id, slug, label, description, prompt_template, is_default, created_at, updated_at
		FROM personas
		WHERE id = $1
	`
	persona := &model.Persona{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&persona.ID, &persona.UserID, &persona.Slug, &persona.Label,
		&persona.Description, &persona.PromptTemplate, &persona.IsDefault,
		&persona.CreatedAt, &persona.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrNotFound
		}
		return nil, err
	}
	return persona, nil
}

// GetDefault retrieves the default persona (is_default = TRUE and user_id IS NULL)
func (s *PersonasStore) GetDefault(ctx context.Context) (*model.Persona, error) {
	query := `
		SELECT id, user_id, slug, label, description, prompt_template, is_default, created_at, updated_at
		FROM personas
		WHERE user_id IS NULL AND is_default = TRUE
		LIMIT 1
	`
	persona := &model.Persona{}
	err := s.db.QueryRowContext(ctx, query).Scan(
		&persona.ID, &persona.UserID, &persona.Slug, &persona.Label,
		&persona.Description, &persona.PromptTemplate, &persona.IsDefault,
		&persona.CreatedAt, &persona.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrNotFound
		}
		return nil, err
	}
	return persona, nil
}
