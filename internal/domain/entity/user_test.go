package entity_test

import (
	"testing"

	"github.com/Figaarillo/golerplate/internal/domain/entity"
)

func TestNewUser(t *testing.T) {
	payload := entity.User{
		Email:     "test@example.com",
		Password:  "password123",
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}
	user, err := entity.NewUser(payload)
	if err != nil {
		t.Errorf("Error creating user: %v", err)
	}

	if user == nil {
		t.Error("User is nil")
	}

	if user.Email != payload.Email {
		t.Errorf("Expected email %s, got %s", payload.Email, user.Email)
	}

	if user.FirstName != payload.FirstName {
		t.Errorf("Expected first name %s, got %s", payload.FirstName, user.FirstName)
	}

	if user.LastName != payload.LastName {
		t.Errorf("Expected last name %s, got %s", payload.LastName, user.LastName)
	}

	if user.Age != payload.Age {
		t.Errorf("Expected age %d, got %d", payload.Age, user.Age)
	}

	if user.ID.String() == "" {
		t.Error("User ID is empty")
	}

	if user.CreatedAt.IsZero() {
		t.Error("User CreatedAt is zero")
	}

	if user.UpdatedAt.IsZero() {
		t.Error("User UpdatedAt is zero")
	}

	// Test password hashing
	if err := user.ComparePassword(payload.Password); err != nil {
		t.Errorf("Error comparing passwords: %v", err)
	}
}

func TestUser_Validate(t *testing.T) {
	tests := []struct {
		name        string
		user        entity.User
		expectError bool
	}{
		{
			name:        "Valid User",
			user:        entity.User{Email: "test@example.com", Password: "password123", FirstName: "John", LastName: "Doe", Age: 30},
			expectError: false,
		},
		{
			name:        "Empty Email",
			user:        entity.User{Email: "", Password: "password123", FirstName: "John", LastName: "Doe", Age: 30},
			expectError: true,
		},
		{
			name:        "Empty Password",
			user:        entity.User{Email: "test@example.com", Password: "", FirstName: "John", LastName: "Doe", Age: 30},
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.user.Validate()
			if (err != nil) != test.expectError {
				t.Errorf("Expected error: %v, got: %v", test.expectError, err)
			}
		})
	}
}
