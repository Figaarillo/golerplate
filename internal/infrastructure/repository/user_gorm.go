package repository

import (
	"github.com/Figaarillo/golerplate/internal/domain/entity"
	"gorm.io/gorm"
)

type UserGorm struct {
	db *gorm.DB
}

func NewUserGorm(db *gorm.DB) *UserGorm {
	return &UserGorm{db: db}
}

func (u *UserGorm) ListAll(offset, limit int) ([]entity.User, error) {
	var user []entity.User

	if result := u.db.Model(&entity.User{}).
		Offset(offset).Limit(limit).
		Find(&user); result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (u *UserGorm) GetByID(id entity.ID) (entity.User, error) {
	var user entity.User

	if result := u.db.Preload("Order").First(&user, "id = ?", id); result.Error != nil {
		return entity.User{}, result.Error
	}

	return user, nil
}

func (u *UserGorm) GetByProp(prop string, value interface{}) (entity.User, error) {
	var user entity.User

	if result := u.db.Where(prop+" = ?", value).First(&user); result.Error != nil {
		return entity.User{}, result.Error
	}

	return user, nil
}

func (u *UserGorm) Create(user *entity.User) (entity.User, error) {
	if result := u.db.Create(user); result.Error != nil {
		return entity.User{}, result.Error
	}

	return *user, nil
}

func (u *UserGorm) Update(id entity.ID, payload entity.User) (entity.User, error) {
	var user entity.User

	if result := u.db.First(&user, "id = ?", id); result.Error != nil {
		return entity.User{}, result.Error
	}

	if err := user.Update(payload); err != nil {
		return entity.User{}, err
	}

	if result := u.db.Save(&user); result.Error != nil {
		return entity.User{}, result.Error
	}

	return user, nil
}

func (u *UserGorm) Delete(id entity.ID) error {
	var user entity.User

	if result := u.db.First(&user, "id = ?", id); result.Error != nil {
		return result.Error
	}

	if result := u.db.Delete(&user); result.Error != nil {
		return result.Error
	}

	return nil
}
