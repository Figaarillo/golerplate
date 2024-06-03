package test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/Figaarillo/golerplate/internal/domain/entity"
	"github.com/Figaarillo/golerplate/internal/setup"
)

var category1, category2 *entity.Category

func initCategories() {
	category1, _ = entity.NewCategory(entity.Category{Name: "Test 1", Description: "Test category"})
	category2, _ = entity.NewCategory(entity.Category{Name: "Test 2", Description: "Test category"})
}

func populateCategoryTable(t *testing.T) {
	ClearCategoryTable(t)

	categories := []entity.Category{*category1, *category2}
	for _, category := range categories {
		if err := DB.Create(&category).Error; err != nil {
			t.Fatalf("Could not create test category: %v", err)
		}
	}
}

func ClearCategoryTable(t *testing.T) {
	ClearAllTables(t, []string{"products", "categories"})
}

func assertCategoryCount(t *testing.T, expectedCount int64) {
	var count int64
	DB.Model(&entity.Category{}).Count(&count)
	AssertEqual(t, "number of categories", expectedCount, count)
}

func assertCategoryAttributes(t *testing.T, id, name, description string) {
	var category entity.Category
	DB.First(&category, "id = ?", id)
	AssertEqual(t, "category name", name, category.Name)
	AssertEqual(t, "category description", description, category.Description)
}

func TestCategory(t *testing.T) {
	InitDBAndRouter(t)
	setup.NewCategory(Router, DB)
	initCategories()

	t.Run("Test CRUD operations", func(t *testing.T) {
		populateCategoryTable(t)
		t.Run("ListAllCategories", testListAllCategories)
		t.Run("GetExistingCategory", testGetExistingCategory)
		t.Run("GetNonExistingCategory", testGetNonExistingCategory)
		t.Run("CreateCategory", testCreateNewCategory)
		t.Run("UpdateCategory", testUpdateCategory)
		t.Run("DeleteCategory", testDeleteCategory)
	})

	t.Run("Test non-CRUD operations", func(t *testing.T) {
		populateCategoryTable(t)
		t.Run("CreateExistingCategory", testCreateExistingCategory)
		t.Run("CreateEmptyCategory", testCreateEmptyCategory)
		t.Run("DeleteNonExistingCategory", testDeleteNonExistingCategory)
		t.Run("UpdateWithInvalidProperties", testUpdateCategoryWithInvalidProperties)
	})
}

func testListAllCategories(t *testing.T) {
	req := NewHTTPRequest(t, http.MethodGet, "/api/categories?offset=0&limit=2", nil)
	res := ExecuteHTTPRequest(req)
	categories := ParseResponseData[entity.Category](t, res)

	AssertResponseCode(t, http.StatusOK, res.Code)
	AssertEqual(t, "number of categories", 2, len(categories))
	AssertEqual(t, "first category name", category1.Name, categories[0].Name)
	AssertEqual(t, "second category name", category2.Name, categories[1].Name)
}

func testGetExistingCategory(t *testing.T) {
	req := NewHTTPRequest(t, http.MethodGet, fmt.Sprintf("/api/categories/%s", category1.ID.String()), nil)
	res := ExecuteHTTPRequest(req)
	category := ParseResponseData[entity.Category](t, res)[0]

	AssertResponseCode(t, http.StatusOK, res.Code)
	AssertEqual(t, "category name", category1.Name, category.Name)
}

func testGetNonExistingCategory(t *testing.T) {
	req := NewHTTPRequest(t, http.MethodGet, "/api/categories/91218", nil)
	res := ExecuteHTTPRequest(req)
	AssertResponseCode(t, http.StatusNotFound, res.Code)
}

func testCreateNewCategory(t *testing.T) {
	ClearCategoryTable(t)

	newCategory := entity.Category{Name: "TestCategory", Description: "TestDescription"}
	req := NewHTTPRequest(t, http.MethodPost, "/api/categories", newCategory)
	res := ExecuteHTTPRequest(req)

	AssertResponseCode(t, http.StatusCreated, res.Code)
	assertCategoryCount(t, 1)
}

func testUpdateCategory(t *testing.T) {
	ClearCategoryTable(t)

	newCategory := entity.Category{Name: "TestCategory", Description: "TestDescription"}
	CreateNewEntity(t, "/api/categories", newCategory)

	var createdCategory entity.Category
	DB.First(&createdCategory)

	updatedCategory := entity.Category{Name: "UpdatedTestCategory", Description: "UpdatedTestDescription"}
	req := NewHTTPRequest(t, http.MethodPut, fmt.Sprintf("/api/categories/%s", createdCategory.ID.String()), updatedCategory)
	res := ExecuteHTTPRequest(req)

	AssertResponseCode(t, http.StatusOK, res.Code)
	assertCategoryAttributes(t, createdCategory.ID.String(), "UpdatedTestCategory", "UpdatedTestDescription")
}

func testDeleteCategory(t *testing.T) {
	ClearCategoryTable(t)

	newCategory := entity.Category{Name: "TestCategory", Description: "TestDescription"}
	CreateNewEntity(t, "/api/categories", newCategory)

	var createdCategory entity.Category
	DB.Find(&createdCategory)

	req := NewHTTPRequest(t, http.MethodDelete, fmt.Sprintf("/api/categories/%s", createdCategory.ID.String()), nil)
	res := ExecuteHTTPRequest(req)

	var categories []entity.Category
	DB.Find(&categories)

	AssertResponseCode(t, http.StatusOK, res.Code)
	assertCategoryCount(t, 0)
}

func testCreateExistingCategory(t *testing.T) {
	newCategory, _ := entity.NewCategory(entity.Category{Name: "Test 1", Description: "Test category"})
	req := NewHTTPRequest(t, http.MethodPost, "/api/categories", newCategory)
	res := ExecuteHTTPRequest(req)

	AssertResponseCode(t, http.StatusConflict, res.Code)
}

func testCreateEmptyCategory(t *testing.T) {
	req := NewHTTPRequest(t, http.MethodPost, "/api/categories", nil)
	res := ExecuteHTTPRequest(req)

	AssertResponseCode(t, http.StatusInternalServerError, res.Code)
	AssertEqual(t, "error message", "error: invalid body provided", res.Body.String())
}

func testDeleteNonExistingCategory(t *testing.T) {
	req := NewHTTPRequest(t, http.MethodDelete, "/api/categories/91218", nil)
	res := ExecuteHTTPRequest(req)
	AssertResponseCode(t, http.StatusNotFound, res.Code)
}

func testUpdateCategoryWithInvalidProperties(t *testing.T) {
	var category entity.Category
	DB.Find(&category)

	updatedCategory := `{"name": "", "description": ""}`
	req := NewHTTPRequest(t, http.MethodPut, fmt.Sprintf("/api/categories/%s", category.ID.String()), updatedCategory)
	res := ExecuteHTTPRequest(req)

	AssertResponseCode(t, http.StatusInternalServerError, res.Code)
	AssertEqual(t, "error message", "error: invalid body provided", res.Body.String())
}
