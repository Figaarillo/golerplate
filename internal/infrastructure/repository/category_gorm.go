package repository

import (
	"github.com/Figaarillo/golerplate/internal/domain/entity"
	"gorm.io/gorm"
)

type CategoryGorm struct {
	db *gorm.DB
}

func NewCategoryGorm(db *gorm.DB) *CategoryGorm {
	return &CategoryGorm{db: db}
}

func (c *CategoryGorm) ListAll(offset, limit int) ([]entity.Category, error) {
	var categories []entity.Category

	if result := c.db.Offset(offset).Limit(limit).
		Find(&categories); result.Error != nil {
		return nil, result.Error
	}

	return categories, nil
}

func (c *CategoryGorm) GetByID(id entity.ID) (entity.Category, error) {
	var category entity.Category

	if result := c.db.First(&category, "id = ?", id); result.Error != nil {
		return entity.Category{}, result.Error
	}

	return category, nil
}

func (c *CategoryGorm) GetByName(name string) (entity.Category, error) {
	var category entity.Category

	if result := c.db.First(&category, "name = ?", name); result.Error != nil {
		return entity.Category{}, result.Error
	}

	return category, nil
}

func (c *CategoryGorm) Create(category *entity.Category) (entity.Category, error) {
	if result := c.db.Create(category); result.Error != nil {
		return entity.Category{}, result.Error
	}

	return *category, nil
}

func (c *CategoryGorm) Update(id entity.ID, payload entity.Category) (entity.Category, error) {
	var category entity.Category

	if result := c.db.First(&category, "id = ?", id); result.Error != nil {
		return entity.Category{}, result.Error
	}

	if err := category.Update(payload); err != nil {
		return entity.Category{}, err
	}

	if result := c.db.Save(&category); result.Error != nil {
		return entity.Category{}, result.Error
	}

	return category, nil
}

func (c *CategoryGorm) Delete(id entity.ID) error {
	var category entity.Category

	if result := c.db.First(&category, "id = ?", id); result.Error != nil {
		return result.Error
	}

	if result := c.db.Delete(&category); result.Error != nil {
		return result.Error
	}

	return nil
}
