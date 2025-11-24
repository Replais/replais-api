package service

import (
	"github.com/Replais/replais-api/internal/logger"
	"github.com/Replais/replais-api/internal/store"
)

// Services is an aggregate of all domain services.
type Services struct {
	Users    UsersService
	Contacts ContactsService
	Personas PersonasService
}

// NewServices wires all services with the underlying store implementation and logger.
// Note: We pass logger but not full config - services typically don't need all config.
// If a service needs specific config, pass only that part.
func NewServices(st store.Storage, log logger.Logger) Services {
	return Services{
		Users:    NewUsersService(st.Users(), log),
		Contacts: NewContactsService(st.Contacts(), st.Personas(), log),
		Personas: NewPersonasService(st.Personas(), log),
	}
}
