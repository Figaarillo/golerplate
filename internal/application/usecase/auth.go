package usecase

import (
	"fmt"
	"time"

	"github.com/Figaarillo/golerplate/internal/domain/entity"
	"github.com/Figaarillo/golerplate/internal/domain/repository"
	"github.com/Figaarillo/golerplate/internal/infrastructure/service"
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

func (uc *AuthUseCase) GenerateAccessToken(userID string) (string, error) {
	claims := &entity.Claims{
		UserID:    userID,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Second * time.Duration(uc.env.JWT_EXPIRATION)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(uc.env.JWT_SECRET_KEY))
}

func (uc *AuthUseCase) GenerateRefreshToken(userID string) (string, error) {
	claims := &entity.Claims{
		UserID:    userID,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Second * time.Duration(uc.env.JWT_REFRESH_EXPIRATION)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(uc.env.JWT_SECRET_KEY_REFRESH))
}

func (uc *AuthUseCase) StoreAccessToken(accessToken, refreshToken, userID string) (entity.Token, error) {
	token := entity.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    time.Now().Add(time.Second * time.Duration(uc.env.JWT_EXPIRATION)),
		IssuedAt:     time.Now(),
		TokenType:    "Bearer",
		UserID:       userID,
	}

	if err := uc.repository.StoreToken(token); err != nil {
		return entity.Token{}, fmt.Errorf("could not store access token: %v", err)
	}

	return token, nil
}

func (uc *AuthUseCase) GenerateAndStoreTokens(userId string) (entity.Token, error) {
	var accessToken string
	var refreshToken string
	var err error

	if accessToken, err = uc.GenerateAccessToken(userId); err != nil {
		return entity.Token{}, fmt.Errorf("could not generate access token: %v", err)
	}

	if refreshToken, err = uc.GenerateRefreshToken(userId); err != nil {
		return entity.Token{}, fmt.Errorf("could not generate refresh token: %v", err)
	}

	jwt, err := uc.StoreAccessToken(accessToken, refreshToken, userId)
	if err != nil {
		return entity.Token{}, fmt.Errorf("could not store access token: %v", err)
	}

	return jwt, nil
}

func (uc *AuthUseCase) IsRefreshTokenValid(refreshToken string) (*entity.Claims, error) {
	return service.VerifyJWTAndReturnClaims(refreshToken, []byte(uc.env.JWT_SECRET_KEY_REFRESH))
}

func (uc *AuthUseCase) IsAccessTokenValid(accessToken string) (*entity.Claims, error) {
	return service.VerifyJWTAndReturnClaims(accessToken, []byte(uc.env.JWT_SECRET_KEY))
}

func (uc *AuthUseCase) RefreshAndStoreAccessToken(refreshToken string) (string, error) {
	claims, err := uc.IsRefreshTokenValid(refreshToken)
	if err != nil {
		return "", fmt.Errorf("invalid refresh token: %v", err)
	}

	newAccessToken, err := uc.GenerateAccessToken(claims.UserID)
	if err != nil {
		return "", fmt.Errorf("could not generate access token: %v", err)
	}

	if _, err := uc.StoreAccessToken(newAccessToken, refreshToken, claims.UserID); err != nil {
		return "", fmt.Errorf("could not store access token: %v", err)
	}

	return newAccessToken, nil
}
