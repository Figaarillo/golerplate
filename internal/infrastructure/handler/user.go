package handler

import (
	"net/http"

	"github.com/Figaarillo/golerplate/internal/application/usecase"
	"github.com/Figaarillo/golerplate/internal/domain/entity"
	"github.com/Figaarillo/golerplate/internal/domain/repository"
	"github.com/Figaarillo/golerplate/internal/infrastructure/service"
	"github.com/Figaarillo/golerplate/internal/shared/utils"
)

type UserHandler struct {
	repository repository.UserRepository
	usecase    *usecase.UserUseCase
}

func NewUserHandler(r repository.UserRepository) *UserHandler {
	return &UserHandler{
		repository: r,
		usecase:    usecase.NewUserUseCase(r),
	}
}

// ListAll godoc
// @Summary List all users with pagination
// @Description Get a list of all users with pagination
// @Tags users
// @Accept json
// @Produce json
// @Param offset query int true "Offset"
// @Param limit query int true "Limit"
// @Success 200 {array} entity.User "Users retrieved successfully"
// @Router /api/users [get]
func (h *UserHandler) ListAll(w http.ResponseWriter, r *http.Request) {
	offset, limit, err := utils.GetPagination(r)
	if err != nil {
		utils.NewHTTPResponse(w).BadRequest(err.Error(), nil)
		return
	}

	users, err := h.usecase.ListAll(offset, limit)
	if err != nil {
		utils.NewHTTPResponse(w).NotFound(err.Error(), nil)
		return
	}

	utils.NewHTTPResponse(w).OK("Users retrieved successfully", users)
}

// GetByID godoc
// @Summary Get a user by ID
// @Description Retrieve a user using its ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} entity.User "User retrieved successfully"
// @Router /api/users/{id} [get]
func (h *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetURLParam(r, "id")
	if err != nil {
		utils.NewHTTPResponse(w).BadRequest(err.Error(), nil)
		return
	}

	user, err := h.usecase.GetByID(id)
	if err != nil {
		utils.NewHTTPResponse(w).NotFound(err.Error(), nil)
		return
	}

	utils.NewHTTPResponse(w).OK("User retrieved successfully", user)
}

// Create godoc
// @Summary Create a user
// @Description Create a new user with the provided data
// @Tags users
// @Accept json
// @Produce json
// @Param User body entity.User true "User"
// @Success 201 {object} entity.User "User created successfully"
// @Router /api/users [post]
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	if err := utils.DecodeReqBody(r, &user); err != nil {
		utils.NewHTTPResponse(w).BadRequest(err.Error(), nil)
		return
	}

	if err := service.NewStructValidator(user).Validate(); err != nil {
		utils.NewHTTPResponse(w).BadRequest(err.Error(), nil)
		return
	}

	newUser, err := h.usecase.Create(user)
	if err != nil {
		utils.NewHTTPResponse(w).Conflict(err.Error(), nil)
		return
	}

	utils.NewHTTPResponse(w).Created("User created successfully", map[string]string{"user_id": newUser.ID.String()})
}

// Update godoc
// @Summary Update a user
// @Description Update an existing user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body entity.User true "User data"
// @Success 200 {object} entity.User "User updated successfully"
// @Router /api/users/{id} [put]
func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetURLParam(r, "id")
	if err != nil {
		utils.NewHTTPResponse(w).BadRequest(err.Error(), nil)
		return
	}

	var payload entity.User
	if err := utils.DecodeReqBody(r, &payload); err != nil {
		utils.NewHTTPResponse(w).BadRequest(err.Error(), nil)
		return
	}

	if err := h.usecase.Update(id, payload); err != nil {
		utils.NewHTTPResponse(w).Conflict(err.Error(), nil)
		return
	}

	utils.NewHTTPResponse(w).OK("User updated successfully", map[string]string{"user_id": id})
}

// Delete godoc
// @Summary Delete a user by ID
// @Description Delete an existing user using its ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 "User deleted successfully"
// @Router /api/users/{id} [delete]
func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetURLParam(r, "id")
	if err != nil {
		utils.NewHTTPResponse(w).BadRequest(err.Error(), nil)
		return
	}

	if err := h.usecase.Delete(id); err != nil {
		utils.NewHTTPResponse(w).NotFound(err.Error(), nil)
		return
	}

	utils.NewHTTPResponse(w).OK("User deleted successfully", map[string]string{"user_id": id})
}
