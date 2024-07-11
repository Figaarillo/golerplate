package handler

import (
	"fmt"
	"net/http"

	"github.com/Figaarillo/golerplate/internal/application/usecase"
	"github.com/Figaarillo/golerplate/internal/domain/entity"
	"github.com/Figaarillo/golerplate/internal/domain/repository"
	"github.com/Figaarillo/golerplate/internal/shared/utils"
	"github.com/go-playground/validator/v10"
)

type ClientHandler struct {
	repository repository.ClientRepository
	usecase    *usecase.ClientUseCase
	validator  *validator.Validate
}

func NewClientHandler(r repository.ClientRepository) *ClientHandler {
	return &ClientHandler{
		repository: r,
		usecase:    usecase.NewClientUseCase(r),
		validator:  validator.New(),
	}
}

// ListAll godoc
// @Summary List all clients with pagination
// @Description Get a list of all clients with pagination
// @Tags clients
// @Accept  json
// @Produce  json
// @Param offset query int true "Offset"
// @Param limit query int true "Limit"
// @Success 200 {array} entity.Client "Clients retrieved successfully"
// @Router /api/clients [get]
func (h *ClientHandler) ListAll(w http.ResponseWriter, r *http.Request) {
	offset, limit := utils.GetPagination(r)

	clients, err := h.usecase.ListAll(offset, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	utils.HandleHTTPResponse(w, "Clients retrieved successfully", http.StatusOK, clients)
}

// GetByID godoc
// @Summary Get a client by ID
// @Description Retrieve a client using its ID
// @Tags clients
// @Accept  json
// @Produce  json
// @Param id path int true "Client ID"
// @Success 200 {object} entity.Client "Client retrieved successfully"
// @Router /api/clients/{id} [get]
func (h *ClientHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetURLParam(r, "id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client, err := h.usecase.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	utils.HandleHTTPResponse(w, "Client retrieved successfully", http.StatusOK, client)
}

// Create godoc
// @Summary Create a client
// @Description Create a new client with the provided data
// @Tags clients
// @Accept  json
// @Produce  json
// @Param client body entity.Client true "Client"
// @Success 201 {object} entity.Client "Client created successfully"
// @Router /api/clients [post]
func (h *ClientHandler) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var payload entity.Client
	if err := utils.DecodeReqBody(r, &payload); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if err := h.validator.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		http.Error(w, fmt.Sprintf("validatoin error: %s", errors), http.StatusUnprocessableEntity)
		return
	}

	if err := h.usecase.Create(payload); err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	utils.HandleHTTPResponse(w, "Client created successfully", http.StatusCreated, nil)
}

// Update godoc
// @Summary Update a client
// @Description Update an existing client by ID
// @Tags clients
// @Accept  json
// @Produce  json
// @Param id path int true "Client ID"
// @Param client body entity.Client true "Client data"
// @Success 200 {object} entity.Client "Client updated successfully"
// @Router /api/clients/{id} [put]
func (h *ClientHandler) Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	id, err := utils.GetURLParam(r, "id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var payload entity.Client
	if err := utils.DecodeReqBody(r, &payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.usecase.Update(id, payload); err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	utils.HandleHTTPResponse(w, "Client updated successfully", http.StatusOK, nil)
}

// Delete godoc
// @Summary Delete a client by ID
// @Description Delete an existing client using its ID
// @Tags clients
// @Accept  json
// @Produce  json
// @Param id path int true "Client ID"
// @Success 200 "Client deleted successfully"
// @Router /api/clients/{id} [delete]
func (h *ClientHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetURLParam(r, "id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.usecase.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	utils.HandleHTTPResponse(w, "Client deleted successfully", http.StatusOK, nil)
}
