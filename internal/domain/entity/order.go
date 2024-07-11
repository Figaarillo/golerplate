package entity

import (
	"time"

	"github.com/Figaarillo/golerplate/internal/domain/exeption"
)

type Order struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Status    string    `json:"status" gorm:"not null" validate:"required"`
	Products  []Product `json:"products" gorm:"many2many:order_products;"`
	Client    Client    `json:"client" gorm:"foreignKey:ClientID"`
	Total     float64   `json:"total" gorm:"not null;default:0"`
	ClientID  ID        `json:"client_id" gorm:"not null;type:uuid" validate:"required"`
	ID        ID        `json:"id" gorm:"type:uuid"`
}

func NewOrder(payload Order) (*Order, error) {
	order := &Order{
		ID:       NewID(),
		Total:    calculateTotal(payload.Products),
		Client:   payload.Client,
		Products: payload.Products,
		Status:   "pending",
	}
	return order, nil
}

func (o *Order) Validate() error {
	if o.Total == 0 || len(o.Products) == 0 {
		return exeption.ErrMissingField
	}
	return nil
}

func (o *Order) SetStatus(status string) {
	o.Status = status
}

func calculateTotal(products []Product) float64 {
	total := 0.0
	for _, product := range products {
		total += product.Price * float64(product.Stock)
	}

	return total
}

func (o *Order) AddProduct(product Product, quantity int) {
	if err := product.Validate(); err != nil {
		return
	}

	// if err := product.UpdateStock(quantity); err != nil {
	// 	return
	// }

	o.Products = append(o.Products, product)
}
