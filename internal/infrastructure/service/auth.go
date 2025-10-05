package service

import (
	"fmt"
	"time"

	"github.com/Figaarillo/golerplate/internal/domain/entity"
	"github.com/golang-jwt/jwt"
)

func GenerateJWT(id string, credential map[string]string, key []byte, exp int) (string, error) {
	claims := &entity.Claims{
		ID:         id,
		Credential: credential,
		IssuedAt:   time.Now().Unix(),
		ExpiresAt:  time.Now().Add(time.Second * time.Duration(exp)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}

func CreateToken(ID string, accessToken string, refreshToken string, exp int) entity.Token {
	return entity.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ID:           ID,
	}
}

func VerifyJWTAndReturnClaims(jwtString string, key []byte) (*entity.Claims, error) {
	claims := &entity.Claims{}

	token, err := jwt.ParseWithClaims(jwtString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	if claims.ExpiresAt < time.Now().Unix() {
		return nil, fmt.Errorf("token has expired")
	}

	return claims, nil
}
