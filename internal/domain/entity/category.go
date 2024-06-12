package entity

import (
	"time"

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
	c.validateName()
	c.validateDescription()

	return nil
}

func (c *Category) validateName() error {
	if err := utils.EnsureValueIsNotEmpty(c.Name); err != nil {
		return err
	}

	return nil
}

func (c *Category) validateDescription() error {
	if err := utils.EnsureValueIsNotEmpty(c.Description); err != nil {
		return err
	}

	return nil
}
