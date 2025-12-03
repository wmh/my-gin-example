package controllers

import (
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wmh/my-gin-example/app/core"
	"github.com/wmh/my-gin-example/app/models"
)

func CreateProduct(c *gin.Context) {
	var req models.ProductCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existing models.Product
	if err := core.DB.Where("sku = ?", req.SKU).First(&existing).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Product with this SKU already exists"})
		return
	}

	product := models.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		Category:    req.Category,
		SKU:         req.SKU,
		IsActive:    true,
	}

	if err := core.DB.Create(&product).Error; err != nil {
		core.ErrorLog("create_product", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	core.Log("create_product", core.H{"product_id": product.ID, "sku": product.SKU})
	c.JSON(http.StatusCreated, product)
}

func GetProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	var product models.Product
	if err := core.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func UpdateProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	var req models.ProductUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var product models.Product
	if err := core.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	updates := map[string]interface{}{}
	if req.Name != "" {
		updates["name"] = req.Name
	}
	if req.Description != "" {
		updates["description"] = req.Description
	}
	if req.Price > 0 {
		updates["price"] = req.Price
	}
	if req.Stock >= 0 {
		updates["stock"] = req.Stock
	}
	if req.Category != "" {
		updates["category"] = req.Category
	}
	if req.IsActive != nil {
		updates["is_active"] = *req.IsActive
	}

	if len(updates) > 0 {
		if err := core.DB.Model(&product).Updates(updates).Error; err != nil {
			core.ErrorLog("update_product", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
			return
		}
	}

	if err := core.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch updated product"})
		return
	}

	core.Log("update_product", core.H{"product_id": product.ID})
	c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	var product models.Product
	if err := core.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	if err := core.DB.Delete(&product).Error; err != nil {
		core.ErrorLog("delete_product", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	core.Log("delete_product", core.H{"product_id": id})
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

func ListProducts(c *gin.Context) {
	var query models.PaginationQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if query.Page == 0 {
		query.Page = 1
	}
	if query.PageSize == 0 {
		query.PageSize = 10
	}
	if query.Order == "" {
		query.Order = "desc"
	}
	if query.SortBy == "" {
		query.SortBy = "created_at"
	}

	db := core.DB.Model(&models.Product{})

	if query.Search != "" {
		search := "%" + query.Search + "%"
		db = db.Where("name LIKE ? OR description LIKE ? OR category LIKE ?", search, search, search)
	}

	var total int64
	if err := db.Count(&total).Error; err != nil {
		core.ErrorLog("list_products", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count products"})
		return
	}

	offset := (query.Page - 1) * query.PageSize
	orderClause := query.SortBy + " " + query.Order

	var products []models.Product
	if err := db.Order(orderClause).Limit(query.PageSize).Offset(offset).Find(&products).Error; err != nil {
		core.ErrorLog("list_products", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	totalPages := int(math.Ceil(float64(total) / float64(query.PageSize)))

	response := models.PaginatedResponse{
		Data:       products,
		Page:       query.Page,
		PageSize:   query.PageSize,
		TotalItems: total,
		TotalPages: totalPages,
	}

	c.JSON(http.StatusOK, response)
}
