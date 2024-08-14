package usecase

import (
	"log"
	"time"

	"github.com/Figaarillo/golerplate/internal/domain/entity"
	"github.com/Figaarillo/golerplate/internal/domain/repository"
	"github.com/Figaarillo/golerplate/internal/shared/config"
	"github.com/golang-jwt/jwt"
)

type AuthUseCase struct {
	env        *config.EnvVars
	repository repository.AuthRepository
}

func NewAuthUseCase(r repository.AuthRepository) *AuthUseCase {
	env, err := config.NewEnvConf(".env")
	if err != nil {
		log.Fatalf("could not load config file: %v", err)
	}

	return &AuthUseCase{
		env:        env,
		repository: r,
	}
}
