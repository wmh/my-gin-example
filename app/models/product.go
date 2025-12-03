package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Name        string         `gorm:"size:100;not null" json:"name" binding:"required"`
	Description string         `gorm:"type:text" json:"description"`
	Price       float64        `gorm:"type:decimal(10,2);not null" json:"price" binding:"required,gt=0"`
	Stock       int            `gorm:"default:0" json:"stock"`
	Category    string         `gorm:"size:50" json:"category"`
	SKU         string         `gorm:"uniqueIndex;size:50" json:"sku"`
	IsActive    bool           `gorm:"default:true" json:"is_active"`
}

type ProductCreateRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Stock       int     `json:"stock" binding:"gte=0"`
	Category    string  `json:"category"`
	SKU         string  `json:"sku" binding:"required"`
}

type ProductUpdateRequest struct {
	Name        string  `json:"name" binding:"omitempty"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"omitempty,gt=0"`
	Stock       int     `json:"stock" binding:"omitempty,gte=0"`
	Category    string  `json:"category"`
	IsActive    *bool   `json:"is_active"`
}

type PaginationQuery struct {
	Page     int    `form:"page" binding:"omitempty,min=1"`
	PageSize int    `form:"page_size" binding:"omitempty,min=1,max=100"`
	Search   string `form:"search"`
	SortBy   string `form:"sort_by"`
	Order    string `form:"order" binding:"omitempty,oneof=asc desc"`
}

type PaginatedResponse struct {
	Data       interface{} `json:"data"`
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
	TotalItems int64       `json:"total_items"`
	TotalPages int         `json:"total_pages"`
}
