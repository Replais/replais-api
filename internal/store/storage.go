package store

import (
	"context"
	"errors"
	"time"

	"github.com/Replais/replais-api/internal/model"
)

var (
	ErrNotFound          = errors.New("resource not found")
	ErrConflict          = errors.New("resource already exists")
	QueryTimeoutDuration = time.Second * 5
)

// Users defines the interface for user repository operations
type Users interface {
	Create(ctx context.Context, user *model.User) error
}

// Contacts defines the interface for contact repository operations
type Contacts interface {
	Create(ctx context.Context, contact *model.Contact) error
	GetByKey(ctx context.Context, userID, platform, contactKey string) (*model.Contact, error)
	GetContactSettingsByContactID(ctx context.Context, contactID string) (*model.ContactSetting, error)
}

type Personas interface {
	Create(ctx context.Context, persona *model.Persona) error
	GetAll(ctx context.Context) ([]model.Persona, error)
	GetByID(ctx context.Context, id string) (*model.Persona, error)
	GetDefault(ctx context.Context) (*model.Persona, error)
}

// Storage is the main storage interface that aggregates all repository interfaces
// Different implementations (Postgres, Mongo, etc.) will implement this
type Storage interface {
	Users() Users
	Contacts() Contacts
	Personas() Personas
}
