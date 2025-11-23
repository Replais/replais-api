package service

import (
	"context"

	"github.com/Replais/replais-api/internal/logger"
	"github.com/Replais/replais-api/internal/model"
	"github.com/Replais/replais-api/internal/store"
)

type PersonasService interface {
	Create(ctx context.Context, persona *model.Persona) error
	GetAll(ctx context.Context) ([]model.Persona, error)
}

type personasService struct {
	personas store.Personas
	logger   logger.Logger
}

func NewPersonasService(personas store.Personas, log logger.Logger) PersonasService {
	return &personasService{
		personas: personas,
		logger:   log,
	}
}

func (s *personasService) Create(ctx context.Context, persona *model.Persona) error {
	s.logger.Info("Creating persona: %s for user %s", persona.Label, persona.UserID)
	// place for validation, dedupe, etc.
	if err := s.personas.Create(ctx, persona); err != nil {
		s.logger.Error("Failed to create persona %s: %v", persona.Label, err)
		return err
	}
	s.logger.Info("Successfully created persona: %s", persona.Label)
	return nil
}

func (s *personasService) GetAll(ctx context.Context) ([]model.Persona, error) {
	personas, err := s.personas.GetAll(ctx)
	if err != nil {
		s.logger.Error("Failed to get all personas: %v", err)
		return nil, err
	}
	return personas, nil
}
