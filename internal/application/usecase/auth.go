package usecase

import (
	"fmt"

	"github.com/Figaarillo/golerplate/internal/domain/entity"
	"github.com/Figaarillo/golerplate/internal/infrastructure/service"
	"github.com/Figaarillo/golerplate/internal/shared/config"
)

type AuthUseCase struct {
	env *config.EnvVars
}

func NewAuthUseCase(env *config.EnvVars) *AuthUseCase {
	return &AuthUseCase{
		env: env,
	}
}

func (uc *AuthUseCase) GenerateAccessToken(userID string, email string) (string, error) {
	key := []byte(uc.env.JWT_SECRET_KEY)
	expiration := uc.env.JWT_EXPIRATION
	credential := map[string]string{"email": email}
	claims, err := service.GenerateJWT(userID, credential, key, expiration)
	if err != nil {
		return "", fmt.Errorf("could not generate access token: %v", err)
	}

	return claims, nil
}

func (uc *AuthUseCase) GenerateRefreshToken(userID string, email string) (string, error) {
	key := []byte(uc.env.JWT_SECRET_KEY_REFRESH)
	expiration := uc.env.JWT_REFRESH_EXPIRATION
	credential := map[string]string{"email": email}
	claims, err := service.GenerateJWT(userID, credential, key, expiration)
	if err != nil {
		return "", fmt.Errorf("could not generate refresh token: %v", err)
	}

	return claims, nil
}

func (uc *AuthUseCase) GenerateTokens(userId string, email string) (entity.Token, error) {
	accessToken, err := uc.GenerateAccessToken(userId, email)
	if err != nil {
		return entity.Token{}, err
	}

	refreshToken, err := uc.GenerateRefreshToken(userId, email)
	if err != nil {
		return entity.Token{}, err
	}

	jwt := service.CreateToken(userId, accessToken, refreshToken, uc.env.JWT_EXPIRATION)

	return jwt, nil
}

func (uc *AuthUseCase) IsRefreshTokenValid(refreshToken string) (*entity.Claims, error) {
	claims, err := service.VerifyJWTAndReturnClaims(refreshToken, []byte(uc.env.JWT_SECRET_KEY_REFRESH))
	if err != nil {
		return nil, fmt.Errorf("invalid refresh token: %v", err)
	}

	return claims, nil
}

func (uc *AuthUseCase) IsAccessTokenValid(accessToken string) (*entity.Claims, error) {
	key := []byte(uc.env.JWT_SECRET_KEY)
	claims, err := service.VerifyJWTAndReturnClaims(accessToken, key)
	if err != nil {
		return nil, fmt.Errorf("invalid access token: %v", err)
	}

	return claims, nil
}

func (uc *AuthUseCase) RefreshAndStoreAccessToken(refreshToken string) (string, error) {
	claims, err := uc.IsRefreshTokenValid(refreshToken)
	if err != nil {
		return "", err
	}

	newAccessToken, err := uc.GenerateAccessToken(claims.Id, claims.Credential["email"])
	if err != nil {
		return "", err
	}

	return newAccessToken, nil
}
