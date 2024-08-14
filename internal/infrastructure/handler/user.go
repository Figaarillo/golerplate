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

type UserHandler struct {
	repository repository.UserRepository
	usecase    *usecase.UserUseCase
	validator  *validator.Validate
}

func NewUserHandler(r repository.UserRepository) *UserHandler {
	return &UserHandler{
		repository: r,
		usecase:    usecase.NewUserUseCase(r),
		validator:  validator.New(),
	}
}

// ListAll godoc
// @Summary List all users with pagination
// @Description Get a list of all users with pagination
// @Tags users
// @Accept  json
// @Produce  json
// @Param offset query int true "Offset"
// @Param limit query int true "Limit"
// @Success 200 {array} entity.User "Users retrieved successfully"
// @Router /api/users [get]
func (h *UserHandler) ListAll(w http.ResponseWriter, r *http.Request) {
	offset, limit := utils.GetPagination(r)

	users, err := h.usecase.ListAll(offset, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	utils.HandleHTTPResponse(w, "Users retrieved successfully", http.StatusOK, users)
}

// GetByID godoc
// @Summary Get a user by ID
// @Description Retrieve a user using its ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} entity.User "User retrieved successfully"
// @Router /api/users/{id} [get]
func (h *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetURLParam(r, "id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := h.usecase.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	utils.HandleHTTPResponse(w, "User retrieved successfully", http.StatusOK, user)
}

// Create godoc
// @Summary Create a user
// @Description Create a new user with the provided data
// @Tags users
// @Accept  json
// @Produce  json
// @Param User body entity.User true "User"
// @Success 201 {object} entity.User "User created successfully"
// @Router /api/users [post]
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var payload entity.User
	if err := utils.DecodeReqBody(r, &payload); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if err := h.validator.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		http.Error(w, fmt.Sprintf("validatoin error: %s", errors), http.StatusUnprocessableEntity)
		return
	}

	user, err := h.usecase.Create(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	res := map[string]string{"id": user.ID.String()}
	utils.HandleHTTPResponse(w, "User created successfully", http.StatusCreated, res)
}

// Update godoc
// @Summary Update a user
// @Description Update an existing user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param user body entity.User true "User data"
// @Success 200 {object} entity.User "User updated successfully"
// @Router /api/users/{id} [put]
func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetURLParam(r, "id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var payload entity.User
	if err := utils.DecodeReqBody(r, &payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.usecase.Update(id, payload); err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	utils.HandleHTTPResponse(w, "User updated successfully", http.StatusOK, nil)
}

// Delete godoc
// @Summary Delete a user by ID
// @Description Delete an existing user using its ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 "User deleted successfully"
// @Router /api/users/{id} [delete]
func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetURLParam(r, "id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.usecase.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	utils.HandleHTTPResponse(w, "User deleted successfully", http.StatusOK, nil)
}
