package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/wmh/my-gin-example/app/core"
	"github.com/wmh/my-gin-example/app/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) {
	var err error
	core.DB, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to initialize test database: %v", err)
	}
	if err := core.DB.AutoMigrate(&models.Product{}); err != nil {
		t.Fatalf("Failed to migrate test database: %v", err)
	}
}

func TestCreateProduct(t *testing.T) {
	setupTestDB(t)
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.POST("/products", CreateProduct)

	productReq := models.ProductCreateRequest{
		Name:        "Test Product",
		Description: "Test Description",
		Price:       99.99,
		Stock:       10,
		Category:    "Electronics",
		SKU:         "TEST-001",
	}

	jsonData, _ := json.Marshal(productReq)
	req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response models.Product
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "Test Product", response.Name)
	assert.Equal(t, 99.99, response.Price)
}

func TestListProducts(t *testing.T) {
	setupTestDB(t)
	gin.SetMode(gin.TestMode)

	core.DB.Create(&models.Product{
		Name:     "Product 1",
		Price:    10.00,
		Stock:    5,
		SKU:      "SKU-001",
		IsActive: true,
	})

	r := gin.Default()
	r.GET("/products", ListProducts)

	req, _ := http.NewRequest("GET", "/products?page=1&page_size=10", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response models.PaginatedResponse
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, 1, response.Page)
	assert.GreaterOrEqual(t, response.TotalItems, int64(1))
}
