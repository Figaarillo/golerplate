package handler

import (
	"net/http"

	"github.com/Figaarillo/golerplate/internal/application/usecase"
	"github.com/Figaarillo/golerplate/internal/domain/entity"
	"github.com/Figaarillo/golerplate/internal/domain/repository"
	"github.com/Figaarillo/golerplate/internal/shared/utils"
	"github.com/go-playground/validator/v10"
)

type OrderHandler struct {
	repository repository.OrderRepository
	usecase    *usecase.OrderUseCase
	validator  *validator.Validate
}

func NewOrderHandler(r repository.OrderRepository) *OrderHandler {
	return &OrderHandler{
		repository: r,
		usecase:    usecase.NewOrderUseCase(r),
		validator:  validator.New(),
	}
}

// ListAll godoc
// @Summary List orders with pagination
// @Description Get a list of orders with pagination
// @Tags orders
// @Accept json
// @Produce json
// @Param offset query int true "Offset"
// @Param limit query int true "Limit"
// @Success 200 {array} entity.Order "Orders retrieved successfully"
// @Router /api/orders [get]
func (h *OrderHandler) List(w http.ResponseWriter, r *http.Request) {
	offset, limit := utils.GetPagination(r)

	orders, err := h.usecase.ListAll(offset, limit)
	if err != nil {
		utils.NewHTTPResponse(w).NotFound(err.Error(), nil)
		return
	}

	utils.NewHTTPResponse(w).OK("Orders retrieved successfully", orders)
}

// GetByID godoc
// @Summary Get a order by ID
// @Description Retrieve a order using its ID
// @Tags orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} entity.Order "Order retrieved successfully"
// @Router /api/orders/{id} [get]
func (h *OrderHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetURLParam(r, "id")
	if err != nil {
		utils.NewHTTPResponse(w).InternalServerError(err.Error(), nil)
		return
	}

	order, err := h.usecase.GetByID(id)
	if err != nil {
		utils.NewHTTPResponse(w).NotFound(err.Error(), nil)
		return
	}

	utils.NewHTTPResponse(w).OK("Order retrieved successfully", order)
}

// GetByUserID godoc
// @Summary Get orders by user ID
// @Description Retrieve orders using its user ID
// @Tags orders
// @Produce json
// @Accept json
// @Param id path uuid true "user ID"
// @Success 200 {array} entity.Order "Orders retrieved successfully"
// @Router /api/orders/user/{id} [get]
func (h *OrderHandler) GetByUserID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetURLParam(r, "id")
	if err != nil {
		utils.NewHTTPResponse(w).InternalServerError(err.Error(), nil)
		return
	}

	orders, err := h.usecase.GetByUserID(id)
	if err != nil {
		utils.NewHTTPResponse(w).NotFound(err.Error(), nil)
		return
	}

	utils.NewHTTPResponse(w).OK("Orders retrieved successfully", orders)
}

// Create godoc
// @Summary Create a new order
// @Description Create a new order with the provided data
// @Tags orders
// @Accept json
// @Produce json
// @Param order body entity.Order true "Order data"
// @Success 201 {object} entity.Order "Order created successfully"
// @Router /api/orders [post]
func (h *OrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var order entity.Order
	if err := utils.DecodeReqBody(r, &order); err != nil {
		utils.NewHTTPResponse(w).InternalServerError(err.Error(), nil)
		return
	}

	if err := h.validator.Struct(order); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.NewHTTPResponse(w).BadRequest("Invalid request body", errors)
		return
	}

	if err := h.usecase.Create(order); err != nil {
		utils.NewHTTPResponse(w).Conflict(err.Error(), nil)
		return
	}

	utils.NewHTTPResponse(w).Created("Order created successfully", nil)
}

// SetStatus godoc
// @Summary Set status
// @Description Set status of an order provided its ID
// @Tags orders
// @Accept json
// @Produce json
// @Param id path int true "order ID"
// @Param order body entity.Order true "Order data"
// @Success 200 {object} entity.Order "Order status updated successfully"
// @Router /api/orders/{id} [put]
func (h *OrderHandler) SetStatus(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	id, err := utils.GetURLParam(r, "id")
	if err != nil {
		utils.NewHTTPResponse(w).InternalServerError(err.Error(), nil)
		return
	}

	type status struct {
		Status string `json:"status"`
	}
	var s status
	if err := utils.DecodeReqBody(r, &s); err != nil {
		utils.NewHTTPResponse(w).InternalServerError(err.Error(), nil)
		return
	}

	if err := h.usecase.SetStatus(id, s.Status); err != nil {
		utils.NewHTTPResponse(w).Conflict(err.Error(), nil)
		return
	}

	utils.NewHTTPResponse(w).OK("Order status updated successfully", nil)
}

// Delete godoc
// @Summary Delete a order
// @Description Delete an existing order using its ID
// @Tags orders
// @Accept json
// @Produce json
// @Param id path int true "order ID"
// @Success 200 "Order deleted successfully"
// @Router /api/orders/{id} [delete]
func (h *OrderHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetURLParam(r, "id")
	if err != nil {
		utils.NewHTTPResponse(w).InternalServerError(err.Error(), nil)
		return
	}

	if err := h.usecase.Delete(id); err != nil {
		utils.NewHTTPResponse(w).NotFound(err.Error(), nil)
		return
	}

	utils.NewHTTPResponse(w).OK("Order deleted successfully", nil)
}
