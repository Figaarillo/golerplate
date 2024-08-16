package usecase

import (
	"github.com/Figaarillo/golerplate/internal/domain/entity"
	"github.com/Figaarillo/golerplate/internal/domain/repository"
)

type UserUseCase struct {
	repository repository.UserRepository
}

func NewUserUseCase(r repository.UserRepository) *UserUseCase {
	return &UserUseCase{repository: r}
}

func (uc *UserUseCase) ListAll(offset, limit int) ([]entity.User, error) {
	return uc.repository.ListAll(offset, limit)
}

func (uc *UserUseCase) GetByID(id string) (entity.User, error) {
	idParsed, err := entity.ParseID(id)
	if err != nil {
		return entity.User{}, err
	}

	return uc.repository.GetByID(idParsed)
}

func (uc *UserUseCase) GetByProp(prop string, value interface{}) (entity.User, error) {
	return uc.repository.GetByProp(prop, value)
}

func (uc *UserUseCase) Create(u entity.User) (entity.User, error) {
	user, err := entity.NewUser(u)
	if err != nil {
		return entity.User{}, err
	}

	return uc.repository.Create(user)
}

func (uc *UserUseCase) Update(id string, payload entity.User) error {
	idParsed, err := entity.ParseID(id)
	if err != nil {
		return err
	}

	uc.repository.Update(idParsed, payload)

	return nil
}

func (uc *UserUseCase) Delete(id string) error {
	idParsed, err := entity.ParseID(id)
	if err != nil {
		return err
	}

	return uc.repository.Delete(idParsed)
}
