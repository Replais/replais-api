package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Replais/replais-api/internal/model"
	"github.com/Replais/replais-api/internal/service"
)

type UsersController struct {
	users service.UsersService
}

func NewUserController(s service.UsersService) *UsersController {
	return &UsersController{users: s}
}

// DTO for incoming payload
type createUserRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func (c *UsersController) Create(w http.ResponseWriter, r *http.Request) {
	var req createUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	user := &model.User{
		Email: req.Email,
		Name:  req.Name,
	}

	if err := c.users.Create(r.Context(), user); err != nil {
		http.Error(w, "failed to create user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(user)
}
