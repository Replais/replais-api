package controllers

import (
	"net/http"

	"github.com/Replais/replais-api/internal/http/dto"
	"github.com/Replais/replais-api/internal/logger"
	"github.com/Replais/replais-api/internal/service"
)

type PersonasController struct {
	personas service.PersonasService
	logger   logger.Logger
}

func NewPersonasController(s service.PersonasService, log logger.Logger) *PersonasController {
	return &PersonasController{
		personas: s,
		logger:   log,
	}
}

func (c *PersonasController) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreatePersonaRequest
	if err := readJSON(w, r, &req); err != nil {
		badRequestError(c.logger, w, r, err)
		return
	}

	if err := Validate.Struct(req); err != nil {
		badRequestError(c.logger, w, r, err)
		return
	}

	persona := dto.ToModelPersona(req)

	if err := c.personas.Create(r.Context(), persona); err != nil {
		internalServerError(c.logger, w, r, err)
		return
	}

	// Return the created persona as response DTO
	response := dto.ToPersonaResponse(*persona)
	err := jsonResponse(w, http.StatusCreated, response)
	if err != nil {
		internalServerError(c.logger, w, r, err)
		return
	}
}

func (c *PersonasController) GetAll(w http.ResponseWriter, r *http.Request) {
	// currently only system personas are supported (no user-level personas)
	personas, err := c.personas.GetAll(r.Context())
	if err != nil {
		internalServerError(c.logger, w, r, err)
		return
	}

	response := dto.ToPersonaListResponse(personas)

	err = jsonResponse(w, http.StatusOK, response)
	if err != nil {
		internalServerError(c.logger, w, r, err)
		return
	}
}

func (c *PersonasController) GetAllAdmin(w http.ResponseWriter, r *http.Request) {
	personas, err := c.personas.GetAll(r.Context())
	if err != nil {
		internalServerError(c.logger, w, r, err)
		return
	}

	responses := dto.ToPersonaResponses(personas)
	err = jsonResponse(w, http.StatusOK, responses)
	if err != nil {
		internalServerError(c.logger, w, r, err)
		return
	}
}
