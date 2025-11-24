package dto

import "github.com/Replais/replais-api/internal/model"

func ToContactSettingsResponse(contactSettings model.ContactSetting, persona model.Persona) ContactSettingsResponseWrapper {
	var contactSettingsResponse ContactSettingsResponse = ContactSettingsResponse{
		ID:           contactSettings.ID,
		ContactID:    contactSettings.ContactID,
		Persona:      ToPersonaListItem(persona),
		Instructions: contactSettings.Instructions,
	}
	return ContactSettingsResponseWrapper{
		ContactSettings: contactSettingsResponse,
	}
}
