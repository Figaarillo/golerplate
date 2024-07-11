package usecase

import (
	"github.com/Figaarillo/golerplate/internal/domain/entity"
	"github.com/Figaarillo/golerplate/internal/domain/repository"
)

type ProductUseCase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(r repository.ProductRepository) *ProductUseCase {
	return &ProductUseCase{repository: r}
}

func (uc *ProductUseCase) ListAll(offset, limit int) ([]entity.Product, error) {
	return uc.repository.ListAll(offset, limit)
}

func (uc *ProductUseCase) GetByID(id string) (entity.Product, error) {
	idParsed, err := entity.ParseID(id)
	if err != nil {
		return entity.Product{}, err
	}

	return uc.repository.GetByID(idParsed)
}

func (uc *ProductUseCase) Create(product entity.Product) error {
	p, err := entity.NewProduct(product)
	if err != nil {
		return err
	}

	if _, err := uc.repository.Create(p); err != nil {
		return err
	}

	return nil
}

func (uc *ProductUseCase) Update(id string, payload entity.Product) error {
	idParsed, err := entity.ParseID(id)
	if err != nil {
		return err
	}

	if _, err := uc.repository.Update(idParsed, payload); err != nil {
		return err
	}

	return nil
}

func (uc *ProductUseCase) Delete(id string) error {
	idParsed, err := entity.ParseID(id)
	if err != nil {
		return err
	}

	return uc.repository.Delete(idParsed)
}
