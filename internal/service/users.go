package service

import (
	"context"

	"github.com/Replais/replais-api/internal/model"
	"github.com/Replais/replais-api/internal/store"
)

// UsersService defines business logic for users.
type UsersService interface {
	Create(ctx context.Context, user *model.User) error
	// later: Get, Update, etc.
}

type usersService struct {
	users store.Users
}

func NewUsersService(users store.Users) UsersService {
	return &usersService{users: users}
}

func (s *usersService) Create(ctx context.Context, user *model.User) error {
	// place for validation, defaults, etc.
	return s.users.Create(ctx, user)
}
