package dto

// CreatePersonaRequest represents the request payload for creating a persona
type CreatePersonaRequest struct {
	UserID         *string `json:"user_id"`
	Label          string  `json:"label" validate:"required"`
	Slug           string  `json:"slug" validate:"required"`
	Description    string  `json:"description"`
	PromptTemplate string  `json:"prompt_template" validate:"required"`
	IsDefault      bool    `json:"is_default"`
}

// PersonaListItem represents a persona in public listings (minimal fields)
type PersonaListItem struct {
	Slug  string `json:"slug"`
	Label string `json:"label"`
}

// PersonaResponse represents the full persona response (for admin/internal use)
type PersonaResponse struct {
	ID             string  `json:"id"`
	Label          string  `json:"label"`
	Slug           string  `json:"slug"`
	Description    string  `json:"description"`
	PromptTemplate string  `json:"prompt_template"`
	IsDefault      bool    `json:"is_default"`
	UserID         *string `json:"user_id"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
}

type PersonaListResponse struct {
	Personas []PersonaListItem `json:"personas"`
}

type PersonaListCompleteResponse struct {
	Personas []PersonaResponse `json:"personas"`
}
