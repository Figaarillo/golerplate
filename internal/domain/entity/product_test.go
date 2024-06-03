package entity_test

import (
	"testing"

	"github.com/Figaarillo/golerplate/internal/domain/entity"
	"github.com/google/uuid"
)

func TestNewProduct(t *testing.T) {
	category := entity.Category{Name: "Test Category", Description: "This is a test category"}

	payload := entity.Product{
		Name:        "Test Product",
		Description: "This is a test product",
		Category:    category,
		Stock:       10,
		Price:       25.99,
	}
	product, err := entity.NewProduct(payload)
	if err != nil {
		t.Errorf("Error creating product: %v", err)
	}

	if product == nil {
		t.Error("Product is nil")
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

func TestProduct_Validate(t *testing.T) {
	tests := []struct {
		name        string
		product     entity.Product
		expectError bool
	}{
		{
			name:        "Valid Product",
			product:     entity.Product{Name: "Test Product", Description: "Test Description", Category: entity.Category{Name: "Test Category", Description: "Test Description"}, Stock: 10, Price: 25.99},
			expectError: false,
		},
		{
			name:        "Empty Name",
			product:     entity.Product{Name: "", Description: "Test Description", Category: entity.Category{Name: "Test Category", Description: "Test Description"}, Stock: 10, Price: 25.99},
			expectError: true,
		},
		{
			name:        "Empty Description",
			product:     entity.Product{Name: "Test Product", Description: "", Category: entity.Category{Name: "Test Category", Description: "Test Description"}, Stock: 10, Price: 25.99},
			expectError: true,
		},
		{
			name:        "Empty Category",
			product:     entity.Product{Name: "Test Product", Description: "Test Description", Stock: 10, Price: 25.99, CategoryID: uuid.Nil},
			expectError: true,
		},
		{
			name:        "Zero Stock",
			product:     entity.Product{Name: "Test Product", Description: "Test Description", Category: entity.Category{Name: "Test Category", Description: "Test Description"}, Stock: 0, Price: 25.99},
			expectError: true,
		},
		{
			name:        "Zero Price",
			product:     entity.Product{Name: "Test Product", Description: "Test Description", Category: entity.Category{Name: "Test Category", Description: "Test Description"}, Stock: 10, Price: 0},
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.product.Validate()
			if (err != nil) != test.expectError {
				t.Errorf("Expected error: %v, got: %v", test.expectError, err)
			}
		})
	}
}
