package dto

import (
	"time"

	"github.com/Replais/replais-api/internal/model"
)

func ToPersonaListItem(persona model.Persona) PersonaListItem {
	return PersonaListItem{
		Slug:  persona.Slug,
		Label: persona.Label,
	}
}

func ToPersonaListItems(personas []model.Persona) []PersonaListItem {
	items := make([]PersonaListItem, len(personas))
	for i, persona := range personas {
		items[i] = ToPersonaListItem(persona)
	}
	return items
}

func ToPersonaListResponse(personas []model.Persona) PersonaListResponse {
	return PersonaListResponse{
		Personas: ToPersonaListItems(personas),
	}
}

func ToPersonaResponse(persona model.Persona) PersonaResponse {
	return PersonaResponse{
		ID:             persona.ID,
		Label:          persona.Label,
		Slug:           persona.Slug,
		Description:    persona.Description,
		PromptTemplate: persona.PromptTemplate,
		IsDefault:      persona.IsDefault,
		UserID:         persona.UserID,
		CreatedAt:      persona.CreatedAt.Format(time.RFC3339),
		UpdatedAt:      persona.UpdatedAt.Format(time.RFC3339),
	}
}

func ToPersonaResponses(personas []model.Persona) PersonaListCompleteResponse {
	responses := make([]PersonaResponse, len(personas))
	for i, persona := range personas {
		responses[i] = ToPersonaResponse(persona)
	}
	return PersonaListCompleteResponse{
		Personas: responses,
	}
}

func ToModelPersona(req CreatePersonaRequest) *model.Persona {
	return &model.Persona{
		UserID:         req.UserID,
		Label:          req.Label,
		Slug:           req.Slug,
		Description:    req.Description,
		PromptTemplate: req.PromptTemplate,
		IsDefault:      req.IsDefault,
	}
}
