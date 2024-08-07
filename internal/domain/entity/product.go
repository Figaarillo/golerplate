package entity

import (
	"time"

	"github.com/Figaarillo/golerplate/internal/shared/utils"
)

type Product struct {
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string    `json:"name" gorm:"unique;not null"`
	Description string    `json:"description" gorm:"not null;default:''"`
	Orders      []Order   `json:"-" gorm:"many2many:order_products;"`
	Category    Category  `json:"category" gorm:"foreignKey:CategoryID"`
	Stock       int       `json:"stock" gorm:"default:0" validate:"gte=0"`
	Price       float64   `json:"price" gorm:"default:0" validate:"gte=0"`
	CategoryID  ID        `json:"category_id" gorm:"not null;type:uuid" validate:"required"`
	ID          ID        `json:"id" gorm:"type:uuid"`
}

func NewProduct(payload Product) (*Product, error) {
	product := &Product{
		ID:          NewID(),
		Name:        payload.Name,
		Description: payload.Description,
		CategoryID:  payload.CategoryID,
		Stock:       payload.Stock,
		Price:       payload.Price,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := product.Validate(); err != nil {
		return nil, err
	}

	return product, nil
}

func (p *Product) Update(payload Product) error {
	utils.AssignIfNotEmpty(&p.Name, payload.Name)
	utils.AssignIfNotEmpty(&p.Description, payload.Description)
	utils.AssignIfNonZero(&p.Stock, payload.Stock)
	utils.AssignIfNonZeroFloat(&p.Price, payload.Price)
	utils.AssignUUIDIFNonEmpty(&p.CategoryID, payload.CategoryID)
	p.UpdatedAt = time.Now()

	if err := p.Validate(); err != nil {
		return err
	}

	return nil
}

func (p *Product) Validate() error {
	p.validateName()
	p.validateDescription()
	p.validateStock()
	p.validatePrice()
	p.validateCategoryID()

	return nil
}

func (p *Product) validateName() error {
	if err := utils.EnsureValueIsNotEmpty(p.Name); err != nil {
		return err
	}

	return nil
}

func (p *Product) validateDescription() error {
	if err := utils.EnsureValueIsNotEmpty(p.Description); err != nil {
		return err
	}

	return nil
}

func (p *Product) validateStock() error {
	if err := utils.EnsureNumberValueIsPositive(float64(p.Stock)); err != nil {
		return err
	}

	return nil
}

func (p *Product) validatePrice() error {
	if err := utils.EnsureNumberValueIsPositive(p.Price); err != nil {
		return err
	}

	return nil
}

func (p *Product) validateCategoryID() error {
	if err := utils.EnsureValueIsNotEmpty(p.CategoryID.String()); err != nil {
		return err
	}

	if err := utils.EnsureValueIsAValidUUID(string(p.CategoryID.String())); err != nil {
		return err
	}

	return nil
}
