package entity

import "github.com/golang-jwt/jwt"

type Token struct {
	AccessToken  string
	RefreshToken string
	Id           string
}

type Claims struct {
	Credential map[string]string `json:"credential"`
	jwt.StandardClaims
	Id        string `json:"id"`
	IssuedAt  int64  `json:"iat"`
	ExpiresAt int64  `json:"exp"`
}

type AuthUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=12"`
}
