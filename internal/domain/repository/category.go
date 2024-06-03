package repository

import "github.com/Figaarillo/golerplate/internal/domain/entity"

type CategoryRepository interface {
	ListAll(offset, limit int) ([]entity.Category, error)
	GetByID(id entity.ID) (entity.Category, error)
	GetByName(name string) (entity.Category, error)
	Create(category *entity.Category) (entity.Category, error)
	Update(id entity.ID, payload entity.Category) (entity.Category, error)
	Delete(id entity.ID) error
}
