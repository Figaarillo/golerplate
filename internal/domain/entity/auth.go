package entity

import "time"

type Token struct {
	ExpiresAt    time.Time
	IssuedAt     time.Time
	AccessToken  string
	RefreshToken string
	TokenType    string
	UserID       string
}

type AuthUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=12"`
}
