package model

import "time"

type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Persona struct {
	ID             string    `json:"id"`
	Label          string    `json:"label"`
	Slug           string    `json:"slug"`
	Description    string    `json:"description"`
	PromptTemplate string    `json:"prompt_template"`
	IsDefault      bool      `json:"is_default"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	UserID         *string   `json:"user_id"`
}

type UserSetting struct {
	ID                 string         `json:"id"`
	UserID             string         `json:"user_id"`
	Platform           string         `json:"platform"`
	DefaultPersonaID   string         `json:"default_persona_id"`
	GlobalInstructions string         `json:"global_instructions"`
	Extra              map[string]any `json:"extra"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
}

type Contact struct {
	ID                 string    `json:"id"`
	UserID             string    `json:"user_id"`
	Platform           string    `json:"platform"`
	PlatformContactKey string    `json:"platform_contact_key"`
	DisplayName        string    `json:"display_name"`
	IsGroup            bool      `json:"is_group"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type ContactSetting struct {
	ID           string         `json:"id"`
	UserID       string         `json:"user_id"`
	ContactID    string         `json:"contact_id"`
	PersonaID    string         `json:"persona_id"`
	Instructions string         `json:"instructions"`
	Extra        map[string]any `json:"extra"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

type ReplaiRequest struct {
	ID               string         `json:"id"`
	UserID           string         `json:"user_id"`
	ContactID        string         `json:"contact_id"`
	Platform         string         `json:"platform"`
	Context          map[string]any `json:"context"`
	UserContext      map[string]any `json:"user_context"`
	UIState          map[string]any `json:"ui_state"`
	ModelName        string         `json:"model_name"`
	PromptTokens     int            `json:"prompt_tokens"`
	CompletionTokens int            `json:"completion_tokens"`
	LatencyMs        int            `json:"latency_ms"`
	ReplyText        string         `json:"reply_text"`
	CreatedAt        time.Time      `json:"created_at"`
}
