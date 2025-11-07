package utils

import (
	"github.com/Figaarillo/golerplate/internal/domain/exeption"
	"github.com/google/uuid"
)

func AssignIfNotEmpty(field *string, newValue string) {
	if newValue != "" {
		*field = newValue
	}
}

func AssignIfNonZero(field *int, newValue int) error {
	if newValue != 0 {
		*field = newValue
		return nil
	}
	return exeption.ErrMissingField
}

func AssignIfNonZeroFloat(field *float64, newValue float64) error {
	if newValue != 0 {
		*field = newValue
		return nil
	}

	return exeption.ErrMissingField
}

func AssignUUIDIFNonEmpty(id *uuid.UUID, newID uuid.UUID) {
	if newID != uuid.Nil {
		*id = newID
	}
}
