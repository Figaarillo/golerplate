package repository

import "github.com/Figaarillo/golerplate/internal/domain/entity"

type ClientRepository interface {
	ListAll(offset, limit int) ([]entity.Client, error)
	GetByID(id entity.ID) (entity.Client, error)
	Create(client *entity.Client) (entity.Client, error)
	Update(id entity.ID, payload entity.Client) (entity.Client, error)
	Delete(id entity.ID) error
}
