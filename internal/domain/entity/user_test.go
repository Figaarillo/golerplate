package entity_test

import (
	"testing"
	"time"

	"github.com/Figaarillo/golerplate/internal/domain/entity"
)

var validUserPayload = entity.User{
	Email:     "test@example.com",
	Password:  "password123",
	FirstName: "John",
	LastName:  "Doe",
	Age:       30,
}

func TestNewUser(t *testing.T) {
	user, err := entity.NewUser(validUserPayload)
	if err != nil {
		t.Errorf("Error creating user: %v", err)
	}

	if user == nil {
		t.Error("User is nil")
		return
	}

	if user.ID.String() == "" {
		t.Error("expected non-empty user ID")
	}
}

func TestNewUser_InvalidEmail(t *testing.T) {
	payload := validUserPayload
	payload.Email = "invalid-email"

	user, err := entity.NewUser(payload)
	if err == nil {
		t.Error("expected error for invalid email, got nil")
	}
	if user != nil {
		t.Error("expected user to be nil on error")
	}
}

func TestComparePassword(t *testing.T) {
	payload := validUserPayload
	user, _ := entity.NewUser(payload)

	// Correct password
	if err := user.ComparePassword(payload.Password); err != nil {
		t.Error("expected password to match, got error:", err)
	}

	// Wrong password
	if err := user.ComparePassword("wrong-password"); err == nil {
		t.Error("expected error for wrong password, got nil")
	}
}

func TestValidate_InvalidFields(t *testing.T) {
	tests := []struct {
		name    string
		user    entity.User
		wantErr bool
	}{
		{"Empty email", entity.User{}, true},
		{"Invalid email", entity.User{Email: "invalid@", Password: "Password123!", FirstName: "John", LastName: "Doe", Age: 25}, true},
		{"Empty password", entity.User{Email: "test@example.com"}, true},
		{"Empty first name", entity.User{Email: "test@example.com", Password: "Password123!", LastName: "Doe"}, true},
		{"Empty last name", entity.User{Email: "test@example.com", Password: "Password123!", FirstName: "John"}, true},
		{"Negative age", entity.User{Email: "test@example.com", Password: "Password123!", FirstName: "John", LastName: "Doe", Age: -5}, true},
		{"Too high age", entity.User{Email: "test@example.com", Password: "Password123!", FirstName: "John", LastName: "Doe", Age: 200}, true},
		{"Valid user", validUserPayload, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.user.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("expected error=%v, got %v", tt.wantErr, err)
			}
		})
	}
}

func TestUpdate_PartialFields(t *testing.T) {
	user, _ := entity.NewUser(validUserPayload)

	oldName := user.FirstName
	oldUpdatedAt := user.UpdatedAt

	time.Sleep(1 * time.Millisecond) // ensure UpdatedAt changes

	updatePayload := entity.User{
		LastName: "Smith",
	}
	err := user.Update(updatePayload)
	if err != nil {
		t.Errorf("unexpected error updating user: %v", err)
	}

	if user.FirstName != oldName {
		t.Error("expected first name to remain unchanged")
	}
	if user.LastName != "Smith" {
		t.Errorf("expected last name to be updated, got %s", user.LastName)
	}
	if !user.UpdatedAt.After(oldUpdatedAt) {
		t.Error("expected UpdatedAt to be updated")
	}
}

func TestUserValidate(t *testing.T) {
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
