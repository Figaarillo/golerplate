package entity_test

import (
	"testing"

	"github.com/Figaarillo/golerplate/internal/domain/entity"
)

func TestNewCategory(t *testing.T) {
	payload := entity.Category{
		Name:        "Test Category",
		Description: "This is a test category",
	}
	category, err := entity.NewCategory(payload)
	if err != nil {
		t.Errorf("Error creating category: %v", err)
	}

	if category == nil {
		t.Error("Category is nil")
	}

	if category.Name != payload.Name {
		t.Errorf("Expected name %s, got %s", payload.Name, category.Name)
	}

	if category.Description != payload.Description {
		t.Errorf("Expected description %s, got %s", payload.Description, category.Description)
	}

	if category.ID.String() == "" {
		t.Error("Category ID is empty")
	}

	if category.CreatedAt.IsZero() {
		t.Error("Category CreatedAt is zero")
	}

	if category.UpdatedAt.IsZero() {
		t.Error("Category UpdatedAt is zero")
	}
}

func TestCategory_Validate(t *testing.T) {
	tests := []struct {
		name        string
		category    entity.Category
		expectError bool
	}{
		{
			name:        "Valid Category",
			category:    entity.Category{Name: "Valid Name", Description: "Valid Description"},
			expectError: false,
		},
		{
			name:        "Empty Name",
			category:    entity.Category{Name: "", Description: "Valid Description"},
			expectError: true,
		},
		{
			name:        "Empty Description",
			category:    entity.Category{Name: "Valid Name", Description: ""},
			expectError: true,
		},
		{
			name:        "Empty Name and Description",
			category:    entity.Category{Name: "", Description: ""},
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.category.Validate()
			if (err != nil) != test.expectError {
				t.Errorf("Expected error: %v, got: %v", test.expectError, err)
			}
		})
	}
}
