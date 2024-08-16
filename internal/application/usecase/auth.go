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

	jwt, err := uc.StoreAccessToken(accessToken, refreshToken, userId.String())
	if err != nil {
		return entity.Token{}, fmt.Errorf("could not store access token: %v", err)
	}

	return jwt, nil
}

func (uc *AuthUseCase) validateAndExtractClaims(tokenString string, secretKey []byte) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	if claims.ExpiresAt < time.Now().Unix() {
		return nil, fmt.Errorf("token has expired")
	}

	return claims, nil
}

func (uc *AuthUseCase) IsRefreshTokenValid(refreshToken string) (*Claims, error) {
	return uc.validateAndExtractClaims(refreshToken, []byte(uc.env.JWT_SECRET_KEY_REFRESH))
}

func (uc *AuthUseCase) IsAccessTokenValid(accessToken string) (*Claims, error) {
	return uc.validateAndExtractClaims(accessToken, []byte(uc.env.JWT_SECRET_KEY))
}

func (uc *AuthUseCase) RefreshAndStoreAccessToken(refreshToken string) (string, error) {
	claims, err := uc.IsRefreshTokenValid(refreshToken)
	if err != nil {
		return "", fmt.Errorf("invalid refresh token: %v", err)
	}

	userId, err := entity.ParseID(claims.UserID)
	if err != nil {
		return "", fmt.Errorf("could not parse user id: %v", err)
	}

	newAccessToken, err := uc.GenerateAccessToken(userId)
	if err != nil {
		return "", fmt.Errorf("could not generate access token: %v", err)
	}

	if _, err := uc.StoreAccessToken(newAccessToken, refreshToken, userId.String()); err != nil {
		return "", fmt.Errorf("could not store access token: %v", err)
	}

	return newAccessToken, nil
}
