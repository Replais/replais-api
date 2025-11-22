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
}

// Storage is the main storage interface that aggregates all repository interfaces
// Different implementations (Postgres, Mongo, etc.) will implement this
type Storage interface {
	Users() Users
	Contacts() Contacts
}
