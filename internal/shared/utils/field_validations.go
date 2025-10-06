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

func EnsureNumberValueIsPositive(number float64) error {
	if number <= 0 {
		return exeption.ErrInvalidNumberFieldZero
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
	if len(password) < 8 {
		return exeption.ErrInvalidPassword
	}

	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`\d`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[\W_]`).MatchString(password)

	if !hasUpper || !hasLower || !hasNumber || !hasSpecial {
		return exeption.ErrInvalidPassword
	}

	return nil
}

func EnsureValueIsValidAge(age int) error {
	if age < 0 || age > 120 {
		return exeption.ErrInvalidAge
	}

	return nil
}

func EnsureValueIsAValidUUID(uuid string) error {
	if uuid == "00000000-0000-0000-0000-000000000000" {
		return exeption.ErrInvalidUUIDFormat
	}

	if !regexp.MustCompile(`^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`).MatchString(uuid) {
		return exeption.ErrInvalidUUIDFormat
	}

	return nil
}
