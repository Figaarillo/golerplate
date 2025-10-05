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
		return
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
		name     string
		input    entity.Category
		expected bool
	}{
		{
			name:     "Valid Category",
			input:    entity.Category{Name: "Valid Name", Description: "Valid Description"},
			expected: false,
		},
		{
			name:     "Empty Name",
			input:    entity.Category{Name: "", Description: "Valid Description"},
			expected: true,
		},
		{
			name:     "Empty Description",
			input:    entity.Category{Name: "Valid Name", Description: ""},
			expected: true,
		},
		{
			name:     "Empty Name and Description",
			input:    entity.Category{Name: "", Description: ""},
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.input.Validate()
			if (err != nil) != test.expected {
				t.Errorf("Expected error: %v, got: %v", test.expected, err)
			}
		})
	}
}
