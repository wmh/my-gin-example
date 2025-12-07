package models

import "time"

type FileUpload struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	FileName     string    `json:"file_name" gorm:"not null"`
	OriginalName string    `json:"original_name" gorm:"not null"`
	FilePath     string    `json:"file_path" gorm:"not null"`
	FileSize     int64     `json:"file_size" gorm:"not null"`
	MimeType     string    `json:"mime_type"`
	UploadedBy   uint      `json:"uploaded_by"`
	Category     string    `json:"category"`
	Description  string    `json:"description"`
	IsPublic     bool      `json:"is_public" gorm:"default:false"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type FileUploadRequest struct {
	Category    string `form:"category" binding:"omitempty,oneof=image document video other"`
	Description string `form:"description" binding:"omitempty,max=500"`
	IsPublic    bool   `form:"is_public"`
}

type FileUploadResponse struct {
	ID           uint      `json:"id"`
	FileName     string    `json:"file_name"`
	OriginalName string    `json:"original_name"`
	FileSize     int64     `json:"file_size"`
	MimeType     string    `json:"mime_type"`
	DownloadURL  string    `json:"download_url"`
	CreatedAt    time.Time `json:"created_at"`
}
