package entity

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type Token struct {
	ExpiresAt    time.Time
	IssuedAt     time.Time
	AccessToken  string
	RefreshToken string
	TokenType    string
	UserID       string
}

type Claims struct {
	jwt.StandardClaims
	UserID    string `json:"user_id"`
	IssuedAt  int64  `json:"iat"`
	ExpiresAt int64  `json:"exp"`
}

type AuthUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=12"`
}
