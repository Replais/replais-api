package store

import "context"

// Users defines the interface for user repository operations
type Users interface {
	Create(ctx context.Context) error
}

// Contacts defines the interface for contact repository operations
type Contacts interface {
	Create(ctx context.Context) error
}

// Storage is the main storage interface that aggregates all repository interfaces
// Different implementations (Postgres, Mongo, etc.) will implement this
type Storage struct {
	Users
	Contacts
}
