package entity

import "time"

type Token struct {
	ExpiresAt    time.Time
	IssuedAt     time.Time
	AccessToken  string
	RefreshToken string
	TokenType    string
	UserID       ID
}
