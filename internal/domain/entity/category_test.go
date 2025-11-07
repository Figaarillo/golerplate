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
		return
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

func TestCategory_Update(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		cpayload entity.Category
		// Named input parameters for target function.
		payload entity.Category
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := entity.NewCategory(tt.cpayload)
			if err != nil {
				t.Fatalf("could not construct receiver type: %v", err)
			}
			gotErr := c.Update(tt.payload)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Update() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Update() succeeded unexpectedly")
			}
		})
	}
}
