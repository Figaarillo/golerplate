package utils

import (
	"regexp"

	"github.com/Figaarillo/golerplate/internal/domain/exeption"
)

func EnsureValueIsNotEmpty(field string) error {
	if field == "" {
		return exeption.ErrMissingField
	}

	return nil
}

func EnsureValueIsAValidEmailFormat(email string) error {
	if email[:1] == "@" || email[len(email)-1:] == "@" {
		return exeption.ErrInvalidEmailAddress
	}

	if !regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(email) {
		return exeption.ErrInvalidEmailAddress
	}

	return nil
}

func EnsureValueIsValidPasswordComplexity(password string) error {
	// if !regexp.MustCompile(`^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[a-zA-Z]).{8,}$`).MatchString(password) {
	// 	return exeption.ErrInvalidPassword
	// }

	return nil
}

func EnsureValueIsValidAge(age int) error {
	if age < 0 || age > 120 {
		return exeption.ErrInvalidAge
	}

	return nil
}
