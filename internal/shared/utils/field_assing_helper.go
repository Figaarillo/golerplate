package utils

import "github.com/google/uuid"

func AssignIfNotEmpty(field *string, newValue string) {
	if newValue != "" {
		*field = newValue
	}
}

func AssignIfNonZero(field *int, newValue int) {
	if newValue != 0 {
		*field = newValue
	}
}

func AssignIfNonZeroFloat(field *float64, newValue float64) {
	if newValue != 0 {
		*field = newValue
	}
}

func AssignUUIDIFNonEmpty(id *uuid.UUID, newID uuid.UUID) {
	if newID.String() != "" || newID != uuid.Nil {
		*id = newID
	}
}
