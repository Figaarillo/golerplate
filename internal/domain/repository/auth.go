package repository

import "github.com/Figaarillo/golerplate/internal/domain/entity"

type AuthRepository interface {
	StoreToken(token entity.Token) error
}
