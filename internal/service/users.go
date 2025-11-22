package service

import (
	"context"

	"github.com/Replais/replais-api/internal/logger"
	"github.com/Replais/replais-api/internal/model"
	"github.com/Replais/replais-api/internal/store"
)

type UsersService interface {
	Create(ctx context.Context, user *model.User) error
}

type usersService struct {
	users  store.Users
	logger logger.Logger
}

func NewUsersService(users store.Users, log logger.Logger) UsersService {
	return &usersService{
		users:  users,
		logger: log,
	}
}

func (s *usersService) Create(ctx context.Context, user *model.User) error {
	s.logger.Info("Creating user: %s", user.Email)
	// place for validation, defaults, etc.
	if err := s.users.Create(ctx, user); err != nil {
		return err
	}
	s.logger.Info("Successfully created user: %s", user.Email)
	return nil
}
