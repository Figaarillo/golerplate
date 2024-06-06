package usecase

import (
	"github.com/Figaarillo/golerplate/internal/domain/entity"
	"github.com/Figaarillo/golerplate/internal/domain/repository"
	"github.com/Figaarillo/golerplate/internal/share/exeption"
)

type OrderUseCase struct {
	repository repository.OrderRepository
}

func NewOrderUseCase(r repository.OrderRepository) *OrderUseCase {
	return &OrderUseCase{repository: r}
}

func (uc *OrderUseCase) ListAll(offset, limit int) ([]entity.Order, error) {
	if offset < 0 || limit < 0 || (offset == 0 && limit == 0) {
		return nil, exeption.ErrInvalidURLParams
	}

	return uc.repository.ListAll(offset, limit)
}

func (uc *OrderUseCase) GetByID(id string) (entity.Order, error) {
	idParsed, err := entity.ParseID(id)
	if err != nil {
		return entity.Order{}, err
	}

	return uc.repository.GetByID(idParsed)
}

func (uc *OrderUseCase) GetByClientID(id string) ([]entity.Order, error) {
	idParsed, err := entity.ParseID(id)
	if err != nil {
		return nil, err
	}

	return uc.repository.GetByClientID(idParsed)
}

func (uc *OrderUseCase) Create(order entity.Order) error {
	o, err := entity.NewOrder(order)
	if err != nil {
		return err
	}

	if _, err := uc.repository.Create(o); err != nil {
		return err
	}

	return nil
}

func (uc *OrderUseCase) SetStatus(id string, status string) error {
	idParsed, err := entity.ParseID(id)
	if err != nil {
		return err
	}

	if err := uc.repository.SetStatus(idParsed, status); err != nil {
		return err
	}

	return nil
}

func (uc *OrderUseCase) Delete(id string) error {
	idParsed, err := entity.ParseID(id)
	if err != nil {
		return err
	}

	return uc.repository.Delete(idParsed)
}
