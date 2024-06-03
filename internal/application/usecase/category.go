package usecase

import (
	"github.com/Figaarillo/golerplate/internal/domain/entity"
	"github.com/Figaarillo/golerplate/internal/domain/repository"
	"github.com/Figaarillo/golerplate/internal/share/exeption"
)

type CategoryUseCase struct {
	repository repository.CategoryRepository
}

func NewCategoryUseCase(r repository.CategoryRepository) *CategoryUseCase {
	return &CategoryUseCase{repository: r}
}

func (uc *CategoryUseCase) ListAll(offset, limit int) ([]entity.Category, error) {
	if offset < 0 || limit < 0 || (offset == 0 && limit == 0) {
		return nil, exeption.ErrInvalidURLParams
	}

	return uc.repository.ListAll(offset, limit)
}

func (uc *CategoryUseCase) GetByID(id string) (entity.Category, error) {
	idParsed, err := entity.ParseID(id)
	if err != nil {
		return entity.Category{}, err
	}

	return uc.repository.GetByID(idParsed)
}

func (uc *CategoryUseCase) GetByName(name string) (entity.Category, error) {
	return uc.repository.GetByName(name)
}

func (uc *CategoryUseCase) Create(category entity.Category) error {
	c, err := entity.NewCategory(category)
	if err != nil {
		return err
	}

	if _, err := uc.repository.Create(c); err != nil {
		return err
	}

	return nil
}

func (uc *CategoryUseCase) Update(id string, payload entity.Category) error {
	idParsed, err := entity.ParseID(id)
	if err != nil {
		return err
	}

	if _, err := uc.repository.Update(idParsed, payload); err != nil {
		return err
	}

	return nil
}

func (uc *CategoryUseCase) Delete(id string) error {
	idParsed, err := entity.ParseID(id)
	if err != nil {
		return err
	}

	if err := uc.repository.Delete(idParsed); err != nil {
		return err
	}

	return nil
}
