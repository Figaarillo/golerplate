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

type Claims struct {
	jwt.StandardClaims
	UserID    string `json:"user_id"`
	IssuedAt  int64  `json:"iat"`
	ExpiresAt int64  `json:"exp"`
}

func (uc *AuthUseCase) GenerateAccessToken(userID entity.ID) (string, error) {
	claims := Claims{
		UserID:    userID.String(),
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Second * time.Duration(uc.env.JWT_EXPIRATION)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(uc.env.JWT_SECRET_KEY))
}

func (uc *AuthUseCase) GenerateRefreshToken(userID entity.ID) (string, error) {
	claims := Claims{
		UserID:    userID.String(),
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Second * time.Duration(uc.env.JWT_REFRESH_EXPIRATION)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(uc.env.JWT_SECRET_KEY_REFRESH))
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

func (uc *AuthUseCase) IsRefreshTokenValid(refreshToken string) error {
	var claims Claims
	return validateToken(refreshToken, &claims, []byte(uc.env.JWT_SECRET_KEY_REFRESH))
}

func (uc *AuthUseCase) IsAccessTokenValid(accessToken string) error {
	var claims Claims
	return validateToken(accessToken, &claims, []byte(uc.env.JWT_SECRET_KEY))
}

func validateToken(token string, claims *Claims, key []byte) error {
	parseToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return key, nil
	})

	if !parseToken.Valid {
		return nil
	}

	if claims.ExpiresAt < time.Now().Unix() {
		return nil
	}

	return err
}
