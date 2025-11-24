package dto

type ContactSettingsResponse struct {
	ID           string          `json:"id"`
	ContactID    string          `json:"contact_id"`
	Persona      PersonaListItem `json:"persona"`
	Instructions string          `json:"instructions"`
}

type ContactSettingsResponseWrapper struct {
	ContactSettings ContactSettingsResponse `json:"contact_settings"`
}
