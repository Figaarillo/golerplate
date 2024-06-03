package repository

import (
	"github.com/Figaarillo/golerplate/internal/domain/entity"
	"gorm.io/gorm"
)

type ClientGorm struct {
	db *gorm.DB
}

func NewClientGorm(db *gorm.DB) *ClientGorm {
	return &ClientGorm{db: db}
}

func (c *ClientGorm) ListAll(offset, limit int) ([]entity.Client, error) {
	var clients []entity.Client

	if result := c.db.Model(&entity.Client{}).
		Offset(offset).Limit(limit).
		Find(&clients); result.Error != nil {
		return nil, result.Error
	}

	return clients, nil
}

func (c *ClientGorm) GetByID(id entity.ID) (entity.Client, error) {
	var client entity.Client

	if result := c.db.Preload("Order").First(&client, "id = ?", id); result.Error != nil {
		return entity.Client{}, result.Error
	}

	return client, nil
}

func (c *ClientGorm) Create(client *entity.Client) (entity.Client, error) {
	if result := c.db.Create(client); result.Error != nil {
		return entity.Client{}, result.Error
	}

	return *client, nil
}

func (c *ClientGorm) Update(id entity.ID, payload entity.Client) (entity.Client, error) {
	var client entity.Client

	if result := c.db.First(&client, "id = ?", id); result.Error != nil {
		return entity.Client{}, result.Error
	}

	if err := client.Update(payload); err != nil {
		return entity.Client{}, err
	}

	if result := c.db.Save(&client); result.Error != nil {
		return entity.Client{}, result.Error
	}

	return client, nil
}

func (c *ClientGorm) Delete(id entity.ID) error {
	var client entity.Client

	if result := c.db.First(&client, "id = ?", id); result.Error != nil {
		return result.Error
	}

	if result := c.db.Delete(&client); result.Error != nil {
		return result.Error
	}

	return nil
}
