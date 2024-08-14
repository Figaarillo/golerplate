package usecase

import (
	"fmt"
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

func NewAuthUseCase(env *config.EnvVars, r repository.AuthRepository) *AuthUseCase {
	return &AuthUseCase{
		env:        env,
		repository: r,
	}
}

func (uc *AuthUseCase) GenerateAccessToken(userID entity.ID) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Second * time.Duration(uc.env.JWT_EXPIRATION)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(uc.env.JWT_SECRET_KEY))
}

func (uc *AuthUseCase) GenerateRefreshToken(userID entity.ID) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Second * time.Duration(uc.env.JWT_REFRESH_EXPIRATION)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(uc.env.JWT_SECRET_KEY))
}

func (uc *AuthUseCase) StoreAccessToken(token entity.Token) error {
	return uc.repository.StoreToken(token)
}

func (uc *AuthUseCase) GenerateAndStoreTokens(userId entity.ID) (entity.Token, error) {
	var accessToken string
	var refreshToken string
	var err error

	if accessToken, err = uc.GenerateAccessToken(userId); err != nil {
		return entity.Token{}, fmt.Errorf("could not generate access token: %v", err)
	}

	if refreshToken, err = uc.GenerateRefreshToken(userId); err != nil {
		return entity.Token{}, fmt.Errorf("could not generate refresh token: %v", err)
	}

	jwt := entity.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    time.Now().Add(time.Second * time.Duration(uc.env.JWT_EXPIRATION)),
		IssuedAt:     time.Now(),
		TokenType:    "Bearer",
		UserID:       userId,
	}

	if err = uc.StoreAccessToken(jwt); err != nil {
		return entity.Token{}, fmt.Errorf("could not store access token: %v", err)
	}

	return jwt, nil
}
