package entity_test

import (
	"testing"
	"time"

	"github.com/Figaarillo/golerplate/internal/domain/entity"
)

func newTestProduct(name string, price float64) entity.Product {
	p, _ := entity.NewProduct(entity.Product{
		Name:        name,
		Description: "test",
		CategoryID:  entity.NewID(),
		Stock:       1,
		Price:       price,
	})
	return *p
}

func newTestOrder(clientID entity.ID, items []entity.OrderItem) (*entity.Order, error) {
	return entity.NewOrder(clientID, items)
}

func TestNewOrder_Success(t *testing.T) {
	clientID := entity.NewID()
	product := newTestProduct("P1", 10)

	items := []entity.OrderItem{
		{ProductID: product.ID, Product: product, Quantity: 2, Price: product.Price},
	}

	order, err := newTestOrder(clientID, items)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if order.Total != 20 {
		t.Errorf("Expected total 20, got %v", order.Total)
	}

	if order.Status != entity.OrderStatusPending {
		t.Errorf("Expected status pending, got %v", order.Status)
	}

	if order.CreatedAt.IsZero() || order.UpdatedAt.IsZero() {
		t.Error("Expected timestamps to be set")
	}
}

func TestNewOrder_NoProducts_ShouldFail(t *testing.T) {
	clientID := entity.NewID()

	_, err := newTestOrder(clientID, []entity.OrderItem{})
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestNewOrder_EmptyClientID_ShouldFail(t *testing.T) {
	_, err := newTestOrder(entity.ID{}, []entity.OrderItem{
		{ProductID: entity.NewID(), Quantity: 1, Price: 10},
	})
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestOrder_AddProduct_Success(t *testing.T) {
	clientID := entity.NewID()
	order, _ := newTestOrder(clientID, []entity.OrderItem{
		{ProductID: entity.NewID(), Quantity: 1, Price: 5},
	})

	product := newTestProduct("New", 10)

	if err := order.AddProduct(product, 3); err != nil {
		t.Fatalf("Expected no error adding product, got %v", err)
	}

	expectedTotal := 5 + (10 * 3)
	if order.Total != float64(expectedTotal) {
		t.Errorf("Expected total %v, got %v", expectedTotal, order.Total)
	}
}

func TestOrder_AddProduct_InvalidQuantity_ShouldFail(t *testing.T) {
	clientID := entity.NewID()
	order, _ := newTestOrder(clientID, []entity.OrderItem{
		{ProductID: entity.NewID(), Quantity: 1, Price: 5},
	})

	product := newTestProduct("Test", 10)

	err := order.AddProduct(product, 0)
	if err == nil {
		t.Error("Expected error when adding product with invalid quantity, got nil")
	}
}

func TestOrder_SetStatus_Success(t *testing.T) {
	clientID := entity.NewID()
	order, _ := newTestOrder(clientID, []entity.OrderItem{
		{ProductID: entity.NewID(), Quantity: 1, Price: 5},
	})

	err := order.SetStatus(entity.OrderStatusPaid)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if order.Status != entity.OrderStatusPaid {
		t.Errorf("Expected status %v, got %v", entity.OrderStatusPaid, order.Status)
	}
}

func TestOrder_SetStatus_Invalid_ShouldFail(t *testing.T) {
	clientID := entity.NewID()
	order, _ := newTestOrder(clientID, []entity.OrderItem{
		{ProductID: entity.NewID(), Quantity: 1, Price: 5},
	})

	err := order.SetStatus("invalid_status")
	if err == nil {
		t.Error("Expected error for invalid status, got nil")
	}
}

func TestCalculateTotal(t *testing.T) {
	items := []entity.OrderItem{
		{Quantity: 2, Price: 5},
		{Quantity: 1, Price: 10},
	}

	total := 5*2 + 10*1
	if entity.CalculateTotal(items) != float64(total) {
		t.Errorf("Expected total %v, got %v", total, entity.CalculateTotal(items))
	}
}

func TestOrder_UpdatedAt_ChangesOnMutation(t *testing.T) {
	clientID := entity.NewID()
	order, _ := newTestOrder(clientID, []entity.OrderItem{
		{ProductID: entity.NewID(), Quantity: 1, Price: 5},
	})

	prev := order.UpdatedAt

	// Sleep to ensure timestamp difference is noticeable
	time.Sleep(10 * time.Millisecond)

	order.SetStatus(entity.OrderStatusShipped)

	if !order.UpdatedAt.After(prev) {
		t.Errorf("Expected UpdatedAt to be updated, but it was not")
	}
}
