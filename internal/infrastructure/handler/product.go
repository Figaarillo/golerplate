package handler

import (
	"net/http"

	"github.com/Figaarillo/golerplate/internal/application/usecase"
	"github.com/Figaarillo/golerplate/internal/domain/entity"
	"github.com/Figaarillo/golerplate/internal/domain/repository"
	"github.com/Figaarillo/golerplate/internal/share/utils"
	"github.com/go-playground/validator/v10"
)

type ProductHandler struct {
	repository repository.ProductRepository
	usecase    *usecase.ProductUseCase
	validator  *validator.Validate
}

func NewProductHandler(r repository.ProductRepository) *ProductHandler {
	return &ProductHandler{
		repository: r,
		usecase:    usecase.NewProductUseCase(r),
		validator:  validator.New(),
	}
}

// @Summary List products with pagination
// @Description Get a list of products with pagination
// @Tags products
// @Accept  json
// @Produce  json
// @Param offset query int true "Offset"
// @Param limit query int true "Limit"
// @Success 200 {array} entity.Product
// @Router /api/products [get]
func (h *ProductHandler) List(w http.ResponseWriter, r *http.Request) {
	offset, limit := utils.GetPagination(r)

	products, err := h.usecase.List(offset, limit)
	if err != nil {
		utils.HandleHTTPError(w, err, http.StatusNotFound)
		return
	}

	utils.HandleHTTPResponse(w, "Products retrieved successfully", http.StatusOK, products)
}

// @Summary Get a product
// @Description Get a product by ID
// @Tags products
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Success 200 {object} entity.Product
// @Router /api/products/{id} [get]
func (h *ProductHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetURLParam(r, "id")
	if err != nil {
		utils.HandleHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	product, err := h.usecase.GetByID(id)
	if err != nil {
		utils.HandleHTTPError(w, err, http.StatusNotFound)
		return
	}

	utils.HandleHTTPResponse(w, "Product retrieved successfully", http.StatusOK, product)
}

// @Summary Create a product
// @Description Create a new product
// @Tags products
// @Accept  json
// @Produce  json
// @Param product body entity.Product true "Product"
// @Success 201
// @Router /api/products [post]
func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var product entity.Product
	if err := utils.DecodeReqBody(r, &product); err != nil {
		utils.HandleHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	if err := h.validator.Struct(product); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.HandleHTTPError(w, errors, http.StatusUnprocessableEntity)
		return
	}

	if err := h.usecase.Create(product); err != nil {
		utils.HandleHTTPError(w, err, http.StatusConflict)
		return
	}

	utils.HandleHTTPResponse(w, "Product created successfully", http.StatusCreated, nil)
}

// @Summary Update a product
// @Description Update a product
// @Tags products
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Param product body entity.Product true "Product"
// @Success 200
// @Router /api/products/{id} [put]
func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	id, err := utils.GetURLParam(r, "id")
	if err != nil {
		utils.HandleHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	var payload entity.Product
	if err := utils.DecodeReqBody(r, &payload); err != nil {
		utils.HandleHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	if err := h.usecase.Update(id, payload); err != nil {
		utils.HandleHTTPError(w, err, http.StatusConflict)
		return
	}

	utils.HandleHTTPResponse(w, "Product updated successfully", http.StatusOK, nil)
}

// @Summary Delete a product
// @Description Delete a product
// @Tags products
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Success 200
// @Router /api/products/{id} [delete]
func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetURLParam(r, "id")
	if err != nil {
		utils.HandleHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	if err := h.usecase.Delete(id); err != nil {
		utils.HandleHTTPError(w, err, http.StatusNotFound)
		return
	}

	utils.HandleHTTPResponse(w, "Product deleted successfully", http.StatusOK, nil)
}
