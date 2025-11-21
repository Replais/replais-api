package service

import "github.com/Replais/replais-api/internal/store"

// Services is an aggregate of all domain services.
type Services struct {
	Users    UsersService
	Contacts ContactsService
}

// NewServices wires all services with the underlying store implementation.
func NewServices(st store.Storage) Services {
	return Services{
		Users:    NewUsersService(st.Users()),
		Contacts: NewContactsService(st.Contacts()),
	}
}
