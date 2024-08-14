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

func (uc *AuthUseCase) GenerateAccessToken(userID entity.ID) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Second * time.Duration(uc.env.JWT_EXPIRATION))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(uc.env.JWT_SECRET_KEY))
}

func (uc *AuthUseCase) GenerateRefreshToken(userID entity.ID) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Second * time.Duration(uc.env.JWT_REFRESH_EXPIRATION))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(uc.env.JWT_SECRET_KEY))
}
