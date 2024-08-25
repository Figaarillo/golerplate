package handler

import (
	"net/http"

	"github.com/Figaarillo/golerplate/internal/application/usecase"
	"github.com/Figaarillo/golerplate/internal/domain/entity"
	"github.com/Figaarillo/golerplate/internal/domain/repository"
	"github.com/Figaarillo/golerplate/internal/infrastructure/service"
	"github.com/Figaarillo/golerplate/internal/shared/utils"
)

type ProductHandler struct {
	repository repository.ProductRepository
	usecase    *usecase.ProductUseCase
}

func NewProductHandler(r repository.ProductRepository) *ProductHandler {
	return &ProductHandler{
		repository: r,
		usecase:    usecase.NewProductUseCase(r),
	}
}

// ListAll godoc
// @Summary List all products with pagination
// @Description Get a list of all products with pagination
// @Tags products
// @Accept json
// @Produce json
// @Param offset query int true "Offset"
// @Param limit query int true "Limit"
// @Success 200 {array} entity.Product "Products retrieved successfully"
// @Router /api/products [get]
func (h *ProductHandler) ListAll(w http.ResponseWriter, r *http.Request) {
	offset, limit, err := utils.GetPagination(r)
	if err != nil {
		utils.NewHTTPResponse(w).BadRequest(err.Error(), nil)
		return
	}

	products, err := h.usecase.ListAll(offset, limit)
	if err != nil {
		utils.NewHTTPResponse(w).NotFound(err.Error(), nil)
		return
	}

	utils.NewHTTPResponse(w).OK("Products retrieved successfully", products)
}

// GetByID godoc
// @Summary Get a product by ID
// @Description Retrive a product using its ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} entity.Product "Product retrieved successfully"
// @Router /api/products/{id} [get]
func (h *ProductHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetURLParam(r, "id")
	if err != nil {
		utils.NewHTTPResponse(w).InternalServerError(err.Error(), nil)
		return
	}

	product, err := h.usecase.GetByID(id)
	if err != nil {
		utils.NewHTTPResponse(w).NotFound(err.Error(), nil)
		return
	}

	utils.NewHTTPResponse(w).OK("Product retrieved successfully", product)
}

// Create godoc
// @Summary Create a new product
// @Description Create a new product with provided data
// @Tags products
// @Accept json
// @Produce json
// @Param product body entity.Product true "Product data"
// @Success 201 {object} entity.Product "Product created successfully"
// @Router /api/products [post]
func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var product entity.Product
	if err := utils.DecodeReqBody(r, &product); err != nil {
		utils.NewHTTPResponse(w).InternalServerError(err.Error(), nil)
		return
	}

	if err := service.NewStructValidator(product).Validate(); err != nil {
		utils.NewHTTPResponse(w).BadRequest(err.Error(), nil)
		return
	}

	if err := h.usecase.Create(product); err != nil {
		utils.NewHTTPResponse(w).Conflict(err.Error(), nil)
		return
	}

	utils.NewHTTPResponse(w).Created("Product created successfully", map[string]string{"product_id": product.ID.String()})
}

// Update godoc
// @Summary Update a product by ID
// @Description Update an existing product by ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body entity.Product true "Product data"
// @Success 200 {object} entity.Product "Product updated successfully"
// @Router /api/products/{id} [put]
func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetURLParam(r, "id")
	if err != nil {
		utils.NewHTTPResponse(w).InternalServerError(err.Error(), nil)
		return
	}

	var payload entity.Product
	if err := utils.DecodeReqBody(r, &payload); err != nil {
		utils.NewHTTPResponse(w).InternalServerError(err.Error(), nil)
		return
	}

	if err := h.usecase.Update(id, payload); err != nil {
		utils.NewHTTPResponse(w).Conflict(err.Error(), nil)
		return
	}

	utils.NewHTTPResponse(w).OK("Product updated successfully", map[string]string{"product_id": id})
}

// Delete godoc
// @Summary Delete a product by ID
// @Description Delete an existing product using its ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} entity.Product "Product deleted successfully"
// @Router /api/products/{id} [delete]
func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetURLParam(r, "id")
	if err != nil {
		utils.NewHTTPResponse(w).InternalServerError(err.Error(), nil)
		return
	}

	if err := h.usecase.Delete(id); err != nil {
		utils.NewHTTPResponse(w).NotFound(err.Error(), nil)
		return
	}

	utils.NewHTTPResponse(w).OK("Product deleted successfully", map[string]string{"product_id": id})
}
