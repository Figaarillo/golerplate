package test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/Figaarillo/golerplate/internal/domain/entity"
	"github.com/Figaarillo/golerplate/internal/setup"
)

var (
	product1, product2 *entity.Product
	category           *entity.Category
)

func initProducts() {
	category, _ = entity.NewCategory(entity.Category{Name: "Test 1", Description: "Test category"})
	product1, _ = entity.NewProduct(entity.Product{
		Name:        "Test Product",
		Description: "This is a test product",
		Stock:       10,
		Price:       25.99,
		CategoryID:  category.ID,
	})
	product2, _ = entity.NewProduct(entity.Product{
		Name:        "Test Product 2",
		Description: "This is a test product",
		Stock:       13,
		Price:       30.99,
		CategoryID:  category.ID,
	})
}

func populateProductTable(t *testing.T) {
	ClearProductTable(t)

	products := []entity.Product{*product1, *product2}
	if err := DB.Create(&category).Error; err != nil {
		t.Fatalf("Could not create test category: %v", err)
	}
	for _, product := range products {
		if err := DB.Create(&product).Error; err != nil {
			t.Fatalf("Could not create test product: %v", err)
		}
	}
}

func ClearProductTable(t *testing.T) {
	ClearAllTables(t, []string{"products", "categories"})
}

func assertProductCount(t *testing.T, expectedCount int64) {
	var count int64
	DB.Model(&entity.Product{}).Count(&count)
	AssertEqual(t, "number of products", expectedCount, count)
}

func assertProductAttributes(t *testing.T, id, name, description string, stock int, price float64) {
	var product entity.Product
	DB.First(&product, "id = ?", id)
	AssertEqual(t, "product name", name, product.Name)
	AssertEqual(t, "product description", description, product.Description)
	AssertEqual(t, "product stock", stock, product.Stock)
	AssertEqual(t, "product price", price, product.Price)
}

func TestProduct(t *testing.T) {
	InitDBAndRouter(t)
	setup.NewProduct(Router, DB)
	initProducts()

	t.Run("Test CRUD operations", func(t *testing.T) {
		populateProductTable(t)
		t.Run("ListAllProducts", testListAllProducts)
		t.Run("GetExistingProduct", testGetExistingProduct)
		t.Run("GetNonExistingProduct", testGetNonExistingProduct)
		t.Run("CreateProduct", testCreateNewProduct)
		t.Run("UpdateProduct", testUpdateProduct)
		t.Run("DeleteProduct", testDeleteProduct)
	})
	t.Run("Test non-CRUD operations", func(t *testing.T) {
		populateProductTable(t)
		t.Run("CreateExistingProduct", testCreateExistingProduct)
		t.Run("CreateEmptyProduct", testCreateEmptyProduct)
		t.Run("DeleteNonExistingProduct", testDeleteNonExistingProduct)
		t.Run("UpdateWithInvalidProperties", testUpdateProductWithInvalidProperties)
	})
}

func testListAllProducts(t *testing.T) {
	req := NewHTTPRequest(t, http.MethodGet, "/api/products?offset=0&limit=2", nil)
	res := ExecuteHTTPRequest(req)
	products := ParseResponseData[entity.Product](t, res)

	AssertResponseCode(t, http.StatusOK, res.Code)
	AssertEqual(t, "number of products", 2, len(products))
	assertProductAttributes(t, products[0].ID.String(), product1.Name, product1.Description, product1.Stock, product1.Price)
	assertProductAttributes(t, products[1].ID.String(), product2.Name, product2.Description, product2.Stock, product2.Price)
}

func testGetExistingProduct(t *testing.T) {
	req := NewHTTPRequest(t, http.MethodGet, fmt.Sprintf("/api/products/%s", product1.ID.String()), nil)
	res := ExecuteHTTPRequest(req)
	Product := ParseResponseData[entity.Product](t, res)[0]

	AssertResponseCode(t, http.StatusOK, res.Code)
	AssertEqual(t, "Product name", product1.Name, Product.Name)
}

func testGetNonExistingProduct(t *testing.T) {
	req := NewHTTPRequest(t, http.MethodGet, fmt.Sprintf("/api/products/%s", "non-existing-id"), nil)
	res := ExecuteHTTPRequest(req)
	AssertResponseCode(t, http.StatusNotFound, res.Code)
}

func testCreateNewProduct(t *testing.T) {
	ClearProductTable(t)

	category, _ := entity.NewCategory(entity.Category{Name: "Test 1", Description: "Test category"})
	if err := DB.Create(&category).Error; err != nil {
		t.Fatalf("Could not create a new category: %v", err)
	}
	newProduct, _ := entity.NewProduct(entity.Product{
		Name:        "Test Product",
		Description: "This is a test product",
		Stock:       10,
		Price:       25.99,
		CategoryID:  category.ID,
	})

	req := NewHTTPRequest(t, http.MethodPost, "/api/products", newProduct)
	res := ExecuteHTTPRequest(req)

	AssertResponseCode(t, http.StatusCreated, res.Code)
	assertProductCount(t, 1)
}

func testUpdateProduct(t *testing.T) {
	ClearProductTable(t)

	category, _ := entity.NewCategory(entity.Category{Name: "Test 1", Description: "Test category"})
	if err := DB.Create(&category).Error; err != nil {
		t.Fatalf("Could not create a new category: %v", err)
	}
	newProduct, _ := entity.NewProduct(entity.Product{
		Name:        "Test Product",
		Description: "This is a test product",
		Stock:       10,
		Price:       25.99,
		CategoryID:  category.ID,
	})
	if err := DB.Create(&newProduct).Error; err != nil {
		t.Fatalf("Could not create a new product: %v", err)
	}

	newProduct.Name = "Updated Product"
	newProduct.Description = "Updated Product Description"
	newProduct.Stock = 20
	newProduct.Price = 30.99

	req := NewHTTPRequest(t, http.MethodPut, fmt.Sprintf("/api/products/%s", newProduct.ID.String()), newProduct)
	res := ExecuteHTTPRequest(req)

	AssertResponseCode(t, http.StatusOK, res.Code)
	assertProductAttributes(t, newProduct.ID.String(), newProduct.Name, newProduct.Description, newProduct.Stock, newProduct.Price)
}

func testDeleteProduct(t *testing.T) {
	ClearProductTable(t)

	category, _ := entity.NewCategory(entity.Category{Name: "Test 1", Description: "Test category"})
	if err := DB.Create(&category).Error; err != nil {
		t.Fatalf("Could not create a new category: %v", err)
	}
	newProduct, _ := entity.NewProduct(entity.Product{
		Name:        "Test Product",
		Description: "This is a test product",
		Stock:       10,
		Price:       25.99,
		CategoryID:  category.ID,
	})
	if err := DB.Create(&newProduct).Error; err != nil {
		t.Fatalf("Could not create a new product: %v", err)
	}

	var createdProduct entity.Product
	DB.Where("id = ?", newProduct.ID).First(&createdProduct)

	req := NewHTTPRequest(t, http.MethodDelete, fmt.Sprintf("/api/products/%s", createdProduct.ID.String()), nil)
	res := ExecuteHTTPRequest(req)

	AssertResponseCode(t, http.StatusOK, res.Code)
	assertProductCount(t, 0)
}

func testCreateExistingProduct(t *testing.T) {
	category, _ = entity.NewCategory(entity.Category{Name: "Test 1", Description: "Test category"})
	newProduct, _ := entity.NewProduct(entity.Product{
		Name:        "Test Product",
		Description: "This is a test product",
		Stock:       10,
		Price:       25.99,
		CategoryID:  category.ID,
	})

	req := NewHTTPRequest(t, http.MethodPost, "/api/products", newProduct)
	res := ExecuteHTTPRequest(req)

	AssertResponseCode(t, http.StatusConflict, res.Code)
}

func testCreateEmptyProduct(t *testing.T) {
	req := NewHTTPRequest(t, http.MethodPost, "/api/products", nil)
	res := ExecuteHTTPRequest(req)

	AssertResponseCode(t, http.StatusInternalServerError, res.Code)
	AssertEqual(t, "error message", "error: invalid body provided", res.Body.String())
}

func testDeleteNonExistingProduct(t *testing.T) {
	req := NewHTTPRequest(t, http.MethodDelete, "/api/products/1234567890", nil)
	res := ExecuteHTTPRequest(req)

	AssertResponseCode(t, http.StatusNotFound, res.Code)
}

func testUpdateProductWithInvalidProperties(t *testing.T) {
	var product entity.Product
	DB.Find(&product)

	updatedProduct := `{"name": "", "description": ""}`
	req := NewHTTPRequest(t, http.MethodPut, fmt.Sprintf("/api/products/%s", product.ID.String()), updatedProduct)
	res := ExecuteHTTPRequest(req)

	AssertResponseCode(t, http.StatusInternalServerError, res.Code)
	AssertEqual(t, "error message", "error: invalid body provided", res.Body.String())
}
