package service

import (
	"fmt"
	"time"

	"github.com/Figaarillo/golerplate/internal/domain/entity"
	"github.com/golang-jwt/jwt"
)

var claims = &entity.Claims{}

func VerifyJWTAndReturnClaims(tokenString string, secretKey []byte) (*entity.Claims, error) {
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
