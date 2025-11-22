package controllers

import (
	"net/http"

	"github.com/Replais/replais-api/internal/config"
	"github.com/Replais/replais-api/internal/logger"
	"github.com/Replais/replais-api/internal/model"
	"github.com/Replais/replais-api/internal/service"
)

type UsersController struct {
	users  service.UsersService
	logger logger.Logger
	config config.Config
}

func NewUserController(s service.UsersService, log logger.Logger, cfg config.Config) *UsersController {
	return &UsersController{
		users:  s,
		logger: log,
		config: cfg,
	}
}

// DTO for incoming payload
type CreateUserPayload struct {
	Email string `json:"email" validate:"required,email"`
	Name  string `json:"name" validate:"required,min=3,max=100"`
}

func (c *UsersController) Create(w http.ResponseWriter, r *http.Request) {
	var userPayload CreateUserPayload
	if err := readJSON(w, r, &userPayload); err != nil {
		badRequestError(c.logger, w, r, err)
		return
	}

	if err := Validate.Struct(userPayload); err != nil {
		badRequestError(c.logger, w, r, err)
		return
	}

	user := &model.User{
		Email: userPayload.Email,
		Name:  userPayload.Name,
	}

	if err := c.users.Create(r.Context(), user); err != nil {
		internalServerError(c.logger, w, r, err)
		return
	}

	err := jsonResponse(w, http.StatusCreated, user)
	if err != nil {
		internalServerError(c.logger, w, r, err)
		return
	}
}
