package repository

import "github.com/Figaarillo/golerplate/internal/domain/entity"

type OrderRepository interface {
	ListAll(offset, limit int) ([]entity.Order, error)
	GetByID(id entity.ID) (entity.Order, error)
	GetByClientID(userID entity.ID) ([]entity.Order, error)
	Create(order *entity.Order) (entity.Order, error)
	SetStatus(id entity.ID, status string) error
	Delete(id entity.ID) error
}
