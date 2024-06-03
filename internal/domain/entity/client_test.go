package entity_test

import (
	"testing"

	"github.com/Figaarillo/golerplate/internal/domain/entity"
)

func TestNewClient(t *testing.T) {
	payload := entity.Client{
		Email:     "test@example.com",
		Password:  "password123",
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}
	client, err := entity.NewClient(payload)
	if err != nil {
		t.Errorf("Error creating client: %v", err)
	}

	if client == nil {
		t.Error("Client is nil")
	}

	if client.Email != payload.Email {
		t.Errorf("Expected email %s, got %s", payload.Email, client.Email)
	}

	if client.FirstName != payload.FirstName {
		t.Errorf("Expected first name %s, got %s", payload.FirstName, client.FirstName)
	}

	if client.LastName != payload.LastName {
		t.Errorf("Expected last name %s, got %s", payload.LastName, client.LastName)
	}

	if client.Age != payload.Age {
		t.Errorf("Expected age %d, got %d", payload.Age, client.Age)
	}

	if client.ID.String() == "" {
		t.Error("Client ID is empty")
	}

	if client.CreatedAt.IsZero() {
		t.Error("Client CreatedAt is zero")
	}

	if client.UpdatedAt.IsZero() {
		t.Error("Client UpdatedAt is zero")
	}

	// Test password hashing
	if err := client.ComparePassword(payload.Password); err != nil {
		t.Errorf("Error comparing passwords: %v", err)
	}
}

func TestClient_Validate(t *testing.T) {
	tests := []struct {
		name        string
		client      entity.Client
		expectError bool
	}{
		{
			name:        "Valid Client",
			client:      entity.Client{Email: "test@example.com", Password: "password123", FirstName: "John", LastName: "Doe", Age: 30},
			expectError: false,
		},
		{
			name:        "Empty Email",
			client:      entity.Client{Email: "", Password: "password123", FirstName: "John", LastName: "Doe", Age: 30},
			expectError: true,
		},
		{
			name:        "Empty Password",
			client:      entity.Client{Email: "test@example.com", Password: "", FirstName: "John", LastName: "Doe", Age: 30},
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.client.Validate()
			if (err != nil) != test.expectError {
				t.Errorf("Expected error: %v, got: %v", test.expectError, err)
			}
		})
	}
}
