package usecase

import (
	"github.com/Figaarillo/golerplate/internal/domain/entity"
	"github.com/Figaarillo/golerplate/internal/domain/repository"
)

type ClientUseCase struct {
	repository repository.ClientRepository
}

func NewClientUseCase(r repository.ClientRepository) *ClientUseCase {
	return &ClientUseCase{repository: r}
}

func (uc *ClientUseCase) ListAll(offset, limit int) ([]entity.Client, error) {
	return uc.repository.ListAll(offset, limit)
}

func (uc *ClientUseCase) GetByID(id string) (entity.Client, error) {
	idParsed, err := entity.ParseID(id)
	if err != nil {
		return entity.Client{}, err
	}

	return uc.repository.GetByID(idParsed)
}

func (uc *ClientUseCase) Create(c entity.Client) error {
	category, err := entity.NewClient(c)
	if err != nil {
		return err
	}

	uc.repository.Create(category)

	return nil
}

func (uc *ClientUseCase) Update(id string, payload entity.Client) error {
	idParsed, err := entity.ParseID(id)
	if err != nil {
		return err
	}

	uc.repository.Update(idParsed, payload)

	return nil
}

func (uc *ClientUseCase) Delete(id string) error {
	idParsed, err := entity.ParseID(id)
	if err != nil {
		return err
	}

	return uc.repository.Delete(idParsed)
}
