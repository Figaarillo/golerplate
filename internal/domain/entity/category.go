package entity

import (
	"time"

	"github.com/Figaarillo/golerplate/internal/domain/exeption"
	"github.com/Figaarillo/golerplate/internal/share/utils"
)

type Category struct {
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string    `json:"name" gorm:"unique"`
	Description string    `json:"description" gorm:"not null;default:''"`
	Products    []Product `json:"products" gorm:"foreignKey:CategoryID"`
	ID          ID        `json:"id"`
}

func NewCategory(payload Category) (*Category, error) {
	category := &Category{
		ID:          NewID(),
		Name:        payload.Name,
		Description: payload.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := category.Validate(); err != nil {
		return nil, err
	}

	return category, nil
}

func (c *Category) Update(payload Category) error {
	utils.AssignIfNotEmpty(&c.Name, payload.Name)
	utils.AssignIfNotEmpty(&c.Description, payload.Description)
	c.UpdatedAt = time.Now()

	if err := c.Validate(); err != nil {
		return err
	}

	return nil
}

func (c *Category) Validate() error {
	if c.Name == "" || c.Description == "" {
		return exeption.ErrMissingField
	}

	return nil
}
