package handler

import (
	"net/http"

	"github.com/Figaarillo/golerplate/internal/application/usecase"
	"github.com/Figaarillo/golerplate/internal/domain/entity"
	"github.com/Figaarillo/golerplate/internal/domain/repository"
	"github.com/Figaarillo/golerplate/internal/share/utils"
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

// @Summary List orders with pagination
// @Description Get a list of orders with pagination
// @Tags orders
// @Accept  json
// @Produce  json
// @Param offset query int true "Offset"
// @Param limit query int true "Limit"
// @Success 200 {array} entity.Product
// @Router /api/orders [get]
func (h *OrderHandler) List(w http.ResponseWriter, r *http.Request) {
	offset, limit := utils.GetPagination(r)

	orders, err := h.usecase.List(offset, limit)
	if err != nil {
		utils.HandleHTTPError(w, err, http.StatusNotFound)
		return
	}

	utils.HandleHTTPResponse(w, "Orders retrieved successfully", http.StatusOK, orders)
}

// @Summary Get a order
// @Description Get a order by ID
// @Tags orders
// @Accept  json
// @Produce  json
// @Param id path int true "order ID"
// @Success 200 {object} entity.order
// @Router /api/orders/{id} [get]
func (h *OrderHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetURLParam(r, "id")
	if err != nil {
		utils.HandleHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	order, err := h.usecase.GetByID(id)
	if err != nil {
		utils.HandleHTTPError(w, err, http.StatusNotFound)
		return
	}

	utils.HandleHTTPResponse(w, "Order retrieved successfully", http.StatusOK, order)
}

// @Summary Get orders by client ID
// @Description Get orders by client ID
// @Tags orders
// @Produce json
// @Param id path uuid true "client ID"
// @Success 200
// @Router /api/orders/client/{id} [get]
func (h *OrderHandler) GetByClientID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetURLParam(r, "id")
	if err != nil {
		utils.HandleHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	orders, err := h.usecase.GetByClientID(id)
	if err != nil {
		utils.HandleHTTPError(w, err, http.StatusNotFound)
		return
	}

	utils.HandleHTTPResponse(w, "Orders retrieved successfully", http.StatusOK, orders)
}

// @Summary Create a order
// @Description Create a new order
// @Tags orders
// @Accept  json
// @Produce  json
// @Param order body entity.Product true "Product"
// @Success 201
// @Router /api/orders [post]
func (h *OrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var order entity.Order
	if err := utils.DecodeReqBody(r, &order); err != nil {
		utils.HandleHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	if err := h.validator.Struct(order); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.HandleHTTPError(w, errors, http.StatusUnprocessableEntity)
		return
	}

	if err := h.usecase.Create(order); err != nil {
		utils.HandleHTTPError(w, err, http.StatusConflict)
		return
	}

	utils.HandleHTTPResponse(w, "Order created successfully", http.StatusCreated, nil)
}

// @Sumanry Set status
// @Description Set status
// @Tags orders
// @Accept  json
// @Produce  json
// @Param id path int true "order ID"
// @Success 200
// @Router /api/orders/{id} [put]
func (h *OrderHandler) SetStatus(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	id, err := utils.GetURLParam(r, "id")
	if err != nil {
		utils.HandleHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	type status struct {
		Status string `json:"status"`
	}
	var s status
	if err := utils.DecodeReqBody(r, &s); err != nil {
		utils.HandleHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	if err := h.usecase.SetStatus(id, s.Status); err != nil {
		utils.HandleHTTPError(w, err, http.StatusConflict)
		return
	}

	utils.HandleHTTPResponse(w, "Order status updated successfully", http.StatusOK, nil)
}

// @Summary Delete a order
// @Description Delete a order
// @Tags orders
// @Accept  json
// @Produce  json
// @Param id path int true "order ID"
// @Success 200
// @Router /api/orders/{id} [delete]
func (h *OrderHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetURLParam(r, "id")
	if err != nil {
		utils.HandleHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	if err := h.usecase.Delete(id); err != nil {
		utils.HandleHTTPError(w, err, http.StatusNotFound)
		return
	}

	utils.HandleHTTPResponse(w, "Order deleted successfully", http.StatusOK, nil)
}
