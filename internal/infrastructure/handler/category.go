package handler

import (
	"net/http"

	"github.com/Figaarillo/golerplate/internal/application/usecase"
	"github.com/Figaarillo/golerplate/internal/domain/entity"
	"github.com/Figaarillo/golerplate/internal/domain/repository"
	"github.com/Figaarillo/golerplate/internal/share/utils"
	"github.com/go-playground/validator/v10"
)

type CategoryHandler struct {
	repository repository.CategoryRepository
	usecase    *usecase.CategoryUseCase
	validator  *validator.Validate
}

func NewCategoryHandler(r repository.CategoryRepository) *CategoryHandler {
	return &CategoryHandler{
		repository: r,
		usecase:    usecase.NewCategoryUseCase(r),
		validator:  validator.New(),
	}
}

// @Summary List all categories
// @Description Get a list of all categories
// @Tags categories
// @Accept  json
// @Produce  json
// @Param offset query int true "Offset"
// @Param limit query int true "Limit"
// @Success 200
// @Router /api/categories [get]
func (h *CategoryHandler) ListAll(w http.ResponseWriter, r *http.Request) {
	offset, limit := utils.GetPagination(r)

	categories, err := h.usecase.ListAll(offset, limit)
	if err != nil {
		utils.HandleHTTPError(w, err, http.StatusNotFound)
		return
	}

	utils.HandleHTTPResponse(w, "Categories retrieved successfully", http.StatusOK, categories)
}

// @Summary Get a category
// @Description Get a category by ID
// @Tags categories
// @Accept  json
// @Produce  json
// @Param id path int true "Category ID"
// @Success 200
// @Router /api/categories/{id} [get]
func (h *CategoryHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetURLParam(r, "id")
	if err != nil {
		utils.HandleHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	category, err := h.usecase.GetByID(id)
	if err != nil {
		utils.HandleHTTPError(w, err, http.StatusNotFound)
		return
	}

	utils.HandleHTTPResponse(w, "Category retrieved successfully", http.StatusOK, category)
}

// @Summary Create a category
// @Description Create a new category
// @Tags categories
// @Accept  json
// @Produce  json
// @Param category body entity.Category true "Category"
// @Success 201
// @Router /api/categories [post]
func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var category entity.Category
	if err := utils.DecodeReqBody(r, &category); err != nil {
		utils.HandleHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	if err := h.validator.Struct(category); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.HandleHTTPError(w, errors, http.StatusUnprocessableEntity)
		return
	}

	if err := h.usecase.Create(category); err != nil {
		utils.HandleHTTPError(w, err, http.StatusConflict)
		return
	}

	utils.HandleHTTPResponse(w, "Category created successfully", http.StatusCreated, nil)
}

// @Summary Update a category
// @Description Update a category
// @Tags categories
// @Accept  json
// @Produce  json
// @Param id path int true "Category ID"
// @Param category body entity.Category true "Category"
// @Success 200
// @Router /api/categories/{id} [put]
func (h *CategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	id, err := utils.GetURLParam(r, "id")
	if err != nil {
		utils.HandleHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	var category entity.Category
	if err := utils.DecodeReqBody(r, &category); err != nil {
		utils.HandleHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	if err := h.usecase.Update(id, category); err != nil {
		utils.HandleHTTPError(w, err, http.StatusConflict)
		return
	}

	utils.HandleHTTPResponse(w, "Category updated successfully", http.StatusOK, nil)
}

// @Summary Delete a category
// @Description Delete a category
// @Tags categories
// @Accept  json
// @Produce  json
// @Param id path int true "Category ID"
// @Success 200
// @Router /api/categories/{id} [delete]
func (h *CategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetURLParam(r, "id")
	if err != nil {
		utils.HandleHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	if err := h.usecase.Delete(id); err != nil {
		utils.HandleHTTPError(w, err, http.StatusNotFound)
		return
	}

	utils.HandleHTTPResponse(w, "Category deleted successfully", http.StatusOK, nil)
}
