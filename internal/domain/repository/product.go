package repository

import "github.com/Figaarillo/golerplate/internal/domain/entity"

type ProductRepository interface {
	List(offset, limit int) ([]entity.Product, error)
	GetByID(id entity.ID) (entity.Product, error)
	Create(product *entity.Product) (entity.Product, error)
	Update(id entity.ID, payload entity.Product) (entity.Product, error)
	Delete(id entity.ID) error
}
