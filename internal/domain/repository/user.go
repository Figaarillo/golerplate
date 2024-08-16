package repository

import "github.com/Figaarillo/golerplate/internal/domain/entity"

type UserRepository interface {
	ListAll(offset, limit int) ([]entity.User, error)
	GetByID(id entity.ID) (entity.User, error)
	GetByProp(prop string, value interface{}) (entity.User, error)
	Create(user *entity.User) (entity.User, error)
	Update(id entity.ID, payload entity.User) (entity.User, error)
	Delete(id entity.ID) error
}
