package entity_test

import (
	"testing"

	"github.com/Figaarillo/golerplate/internal/domain/entity"
	"github.com/google/uuid"
)

var category, _ = entity.NewCategory(entity.Category{Name: "Test Category", Description: "This is a test category"})

func TestNewProduct(t *testing.T) {
	payload := entity.Product{
		Name:        "Test Product",
		Description: "This is a test product",
		Category:    *category,
		CategoryID:  category.ID,
		Stock:       10,
		Price:       25.99,
	}
	product, err := entity.NewProduct(payload)
	if err != nil {
		t.Errorf("Error creating product: %v", err)
	}

	if product == nil {
		t.Error("Product is nil")
		return
	}

	if product.Name != payload.Name {
		t.Errorf("Expected name %s, got %s", payload.Name, product.Name)
	}

	if product.Description != payload.Description {
		t.Errorf("Expected description %s, got %s", payload.Description, product.Description)
	}

	if product.Stock != payload.Stock {
		t.Errorf("Expected Stock %d, got %d", payload.Stock, product.Stock)
	}

	if product.Price != payload.Price {
		t.Errorf("Expected price %f, got %f", payload.Price, product.Price)
	}

	if product.ID.String() == "" {
		t.Error("Product ID is empty")
	}

	if product.CreatedAt.IsZero() {
		t.Error("Product CreatedAt is zero")
	}

	if product.UpdatedAt.IsZero() {
		t.Error("Product UpdatedAt is zero")
	}
}

func TestValidate(t *testing.T) {
	tests := []struct {
		name     string
		input    entity.Product
		expected bool
	}{
		{
			name:     "Valid Product",
			input:    entity.Product{Name: "Test Product", Description: "Test Description", Category: *category, CategoryID: category.ID, Stock: 10, Price: 25.99},
			expected: false,
		},
		{
			name:     "Empty Name",
			input:    entity.Product{Name: "", Description: "Test Description", Category: *category, CategoryID: category.ID, Stock: 10, Price: 25.99},
			expected: true,
		},
		{
			name:     "Empty Description",
			input:    entity.Product{Name: "Test Product", Description: "", Category: *category, CategoryID: category.ID, Stock: 10, Price: 25.99},
			expected: true,
		},
		{
			name:     "Empty Category",
			input:    entity.Product{Name: "Test Product", Description: "Test Description", Stock: 10, Price: 25.99, CategoryID: entity.ID(uuid.Nil)},
			expected: true,
		},
		{
			name:     "Zero Stock",
			input:    entity.Product{Name: "Test Product", Description: "Test Description", Category: *category, CategoryID: category.ID, Stock: 0, Price: 25.99},
			expected: true,
		},
		{
			name:     "Zero Price",
			input:    entity.Product{Name: "Test Product", Description: "Test Description", Category: *category, CategoryID: category.ID, Stock: 10, Price: 0},
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.input.Validate()
			if (err != nil) != test.expected {
				t.Errorf("Expected error: %v, got error: %v", test.expected, err)
			}
		})
	}
}
