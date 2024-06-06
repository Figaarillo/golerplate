package repository

import (
	"github.com/Figaarillo/golerplate/internal/domain/entity"
	"gorm.io/gorm"
)

type OrderGorm struct {
	db *gorm.DB
}

func NewOrderGorm(db *gorm.DB) *OrderGorm {
	return &OrderGorm{db: db}
}

func (o *OrderGorm) ListAll(offset, limit int) ([]entity.Order, error) {
	var orders []entity.Order
	if result := o.db.Model(&entity.Order{}).
		Offset(offset).Limit(limit).
		// Select([]string{"id", "user_id", "total", "status"}).
		Find(&orders); result.Error != nil {
		return nil, result.Error
	}

	return orders, nil
}

func (o *OrderGorm) GetByID(id entity.ID) (entity.Order, error) {
	var order entity.Order
	if result := o.db.
		Preload("Products").Preload("User").
		First(&order, "id = ?", id); result.Error != nil {
		return entity.Order{}, result.Error
	}

	return order, nil
}

func (o *OrderGorm) GetByClientID(userID entity.ID) ([]entity.Order, error) {
	var orders []entity.Order
	if result := o.db.
		Preload("Products").Preload("User").
		Find(&orders, "user_id = ?", userID); result.Error != nil {
		return nil, result.Error
	}

	return orders, nil
}

func (o *OrderGorm) Create(order *entity.Order) (entity.Order, error) {
	if result := o.db.Create(&order); result.Error != nil {
		return entity.Order{}, result.Error
	}

	return *order, nil
}

func (o *OrderGorm) SetStatus(id entity.ID, status string) error {
	var order entity.Order

	if result := o.db.First(&order, "id = ?", id); result.Error != nil {
		return result.Error
	}

	order.SetStatus(status)

	if result := o.db.Save(&order); result.Error != nil {
		return result.Error
	}

	return nil
}

func (o *OrderGorm) Delete(id entity.ID) error {
	var order entity.Order

	if result := o.db.First(&order, "id = ?", id); result.Error != nil {
		return result.Error
	}

	if result := o.db.Delete(&order); result.Error != nil {
		return result.Error
	}

	return nil
}
