package repository

import "github.com/Figaarillo/golerplate/internal/domain/entity"

type AuthRepositoryMemory struct {
	tokens []entity.Token
}

func NewAuthRepositoryMemory() *AuthRepositoryMemory {
	return &AuthRepositoryMemory{}
}

func (r *AuthRepositoryMemory) StoreToken(token entity.Token) error {
	r.tokens = append(r.tokens, token)
	return nil
}
