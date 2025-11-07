package entity

import (
	"time"

	"github.com/Figaarillo/golerplate/internal/domain/exeption"
	"github.com/Figaarillo/golerplate/internal/shared/utils"
)

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusPaid      OrderStatus = "paid"
	OrderStatusShipped   OrderStatus = "shipped"
	OrderStatusCancelled OrderStatus = "cancelled"
)

type Order struct {
	ID        ID          `json:"id" gorm:"type:uuid"`
	ClientID  ID          `json:"client_id" gorm:"not null;type:uuid" validate:"required"`
	Client    Client      `json:"client" gorm:"foreignKey:ClientID"`
	Products  []OrderItem `json:"products" gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE;"`
	Total     float64     `json:"total" gorm:"not null;default:0"`
	Status    OrderStatus `json:"status" gorm:"not null"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

type OrderItem struct {
	OrderID   ID      `json:"-" gorm:"type:uuid"`
	ProductID ID      `json:"product_id" gorm:"not null;type:uuid"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID"`
	Quantity  int     `json:"quantity" gorm:"not null"`
	Price     float64 `json:"price" gorm:"not null"` // price at order time
}

func NewOrder(clientID ID, items []OrderItem) (*Order, error) {
	order := &Order{
		ID:        NewID(),
		ClientID:  clientID,
		Products:  items,
		Total:     CalculateTotal(items),
		Status:    OrderStatusPending,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := order.Validate(); err != nil {
		return nil, err
	}

	return order, nil
}

func (o *Order) Validate() error {
	if err := utils.EnsureValueIsNotEmpty(o.ClientID.String()); err != nil {
		return err
	}

	if len(o.Products) == 0 {
		return exeption.ErrMissingField // mejor: ErrOrderMustHaveProducts
	}

	if err := utils.EnsureNumberValueIsPositive(o.Total); err != nil {
		return err
	}

	if !isValidStatus(o.Status) {
		return exeption.ErrInvalidOrderStatus
	}

	return nil
}

func isValidStatus(s OrderStatus) bool {
	switch s {
	case OrderStatusPending, OrderStatusPaid, OrderStatusShipped, OrderStatusCancelled:
		return true
	}
	return false
}

func CalculateTotal(items []OrderItem) float64 {
	total := 0.0
	for _, item := range items {
		total += item.Price * float64(item.Quantity)
	}
	return total
}

func (o *Order) AddProduct(product Product, quantity int) error {
	if quantity <= 0 {
		return exeption.ErrInvalidQuantity
	}

	if err := product.Validate(); err != nil {
		return err
	}

	item := OrderItem{
		OrderID:   o.ID,
		ProductID: product.ID,
		Product:   product,
		Quantity:  quantity,
		Price:     product.Price,
	}

	o.Products = append(o.Products, item)
	o.Total = CalculateTotal(o.Products)
	o.UpdatedAt = time.Now()
	return nil
}

func (o *Order) SetStatus(status OrderStatus) error {
	if !isValidStatus(status) {
		return exeption.ErrInvalidOrderStatus
	}
	o.Status = status
	o.UpdatedAt = time.Now()
	return nil
}
