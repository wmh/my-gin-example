package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wmh/my-gin-example/app/core"
	"github.com/wmh/my-gin-example/app/models"
	"github.com/wmh/my-gin-example/app/services"
)

var fileService = services.NewFileService()

func UploadFile(c *gin.Context) {
	var req models.FileUploadRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	allowedTypesMap := map[string]string{
		"image":    services.AllowedImageTypes,
		"document": services.AllowedDocTypes,
	}
	allowedTypes := allowedTypesMap[req.Category]

	if err := fileService.ValidateFile(file, allowedTypes); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fileName, filePath, err := fileService.SaveFile(file)
	if err != nil {
		core.ErrorLog("upload_file", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	userID := c.GetUint("user_id")

	fileUpload := models.FileUpload{
		FileName:     fileName,
		OriginalName: file.Filename,
		FilePath:     filePath,
		FileSize:     file.Size,
		MimeType:     file.Header.Get("Content-Type"),
		UploadedBy:   userID,
		Category:     req.Category,
		Description:  req.Description,
		IsPublic:     req.IsPublic,
	}

	if err := core.DB.Create(&fileUpload).Error; err != nil {
		// Log orphaned file for cleanup
		core.ErrorLog("upload_file", fmt.Sprintf("DB error: %s; orphaned file at %s needs cleanup", err.Error(), filePath))
		fileService.DeleteFile(filePath)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file record"})
		return
	}

	response := models.FileUploadResponse{
		ID:           fileUpload.ID,
		FileName:     fileUpload.FileName,
		OriginalName: fileUpload.OriginalName,
		FileSize:     fileUpload.FileSize,
		MimeType:     fileUpload.MimeType,
		DownloadURL:  fmt.Sprintf("/v2/files/%d/download", fileUpload.ID),
		CreatedAt:    fileUpload.CreatedAt,
		Category:     fileUpload.Category,
		Description:  fileUpload.Description,
	}

	core.Log("upload_file", core.H{"file_id": fileUpload.ID, "user_id": userID})
	c.JSON(http.StatusCreated, response)
}

type FileUploadResult struct {
	Success bool                        `json:"success"`
	File    *models.FileUploadResponse  `json:"file,omitempty"`
	Error   string                      `json:"error,omitempty"`
	Name    string                      `json:"name"`
}

func UploadMultipleFiles(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse multipart form"})
		return
	}

	files := form.File["files"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No files uploaded"})
		return
	}

	if len(files) > 10 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Maximum 10 files allowed"})
		return
	}

	// Check total batch size
	var totalSize int64
	for _, file := range files {
		totalSize += file.Size
	}
	if totalSize > services.MaxBatchSize {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":      "Total batch size exceeds limit",
			"total_size": totalSize,
			"max_size":   services.MaxBatchSize,
		})
		return
	}

	userID := c.GetUint("user_id")
	var results []FileUploadResult

	for _, file := range files {
		result := FileUploadResult{
			Name: file.Filename,
		}
		
		if err := fileService.ValidateFile(file, ""); err != nil {
			result.Success = false
			result.Error = err.Error()
			results = append(results, result)
			continue
		}

		fileName, filePath, err := fileService.SaveFile(file)
		if err != nil {
			result.Success = false
			result.Error = fmt.Sprintf("Failed to save: %s", err.Error())
			results = append(results, result)
			continue
		}

		fileUpload := models.FileUpload{
			FileName:     fileName,
			OriginalName: file.Filename,
			FilePath:     filePath,
			FileSize:     file.Size,
			MimeType:     file.Header.Get("Content-Type"),
			UploadedBy:   userID,
			IsPublic:     false,
		}

		if err := core.DB.Create(&fileUpload).Error; err != nil {
			fileService.DeleteFile(filePath)
			result.Success = false
			result.Error = "Failed to save record"
			results = append(results, result)
			continue
		}

		result.Success = true
		result.File = &models.FileUploadResponse{
			ID:           fileUpload.ID,
			FileName:     fileUpload.FileName,
			OriginalName: fileUpload.OriginalName,
			FileSize:     fileUpload.FileSize,
			MimeType:     fileUpload.MimeType,
			DownloadURL:  fmt.Sprintf("/v2/files/%d/download", fileUpload.ID),
			CreatedAt:    fileUpload.CreatedAt,
		}
		results = append(results, result)
	}

	successCount := 0
	for _, r := range results {
		if r.Success {
			successCount++
		}
	}

	core.Log("upload_multiple_files", core.H{"total": len(files), "success": successCount, "user_id": userID})
	c.JSON(http.StatusCreated, gin.H{
		"total":   len(files),
		"success": successCount,
		"failed":  len(files) - successCount,
		"results": results,
	})
}

func GetFile(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file ID"})
		return
	}

	var fileUpload models.FileUpload
	if err := core.DB.First(&fileUpload, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	userID := c.GetUint("user_id")
	if !fileUpload.IsPublic && fileUpload.UploadedBy != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	response := models.FileUploadResponse{
		ID:           fileUpload.ID,
		FileName:     fileUpload.FileName,
		OriginalName: fileUpload.OriginalName,
		FileSize:     fileUpload.FileSize,
		MimeType:     fileUpload.MimeType,
		DownloadURL:  fmt.Sprintf("/v2/files/%d/download", fileUpload.ID),
		CreatedAt:    fileUpload.CreatedAt,
		Category:     fileUpload.Category,
		Description:  fileUpload.Description,
	}

	c.JSON(http.StatusOK, response)
}

func DownloadFile(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file ID"})
		return
	}

	var fileUpload models.FileUpload
	if err := core.DB.First(&fileUpload, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	userID := c.GetUint("user_id")
	if !fileUpload.IsPublic && fileUpload.UploadedBy != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	// Validate file path is within upload directory
	cleanPath := filepath.Clean(fileUpload.FilePath)
	uploadDir := filepath.Clean("./uploads")
	if !strings.HasPrefix(cleanPath, uploadDir) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Invalid file path"})
		return
	}

	// Escape filename for Content-Disposition header
	escapedFilename := strings.ReplaceAll(fileUpload.OriginalName, "\"", "\\\"")
	
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", escapedFilename))
	c.Header("Content-Type", fileUpload.MimeType)
	c.File(cleanPath)
}

func ListFiles(c *gin.Context) {
	var query models.PaginationQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if query.Page == 0 {
		query.Page = 1
	}
	if query.PageSize == 0 {
		query.PageSize = 20
	}

	userID := c.GetUint("user_id")

	db := core.DB.Model(&models.FileUpload{}).Where("uploaded_by = ?", userID)

	var total int64
	if err := db.Count(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count files"})
		return
	}

	offset := (query.Page - 1) * query.PageSize

	var files []models.FileUpload
	if err := db.Order("created_at desc").Limit(query.PageSize).Offset(offset).Find(&files).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch files"})
		return
	}

	var responses []models.FileUploadResponse
	for _, file := range files {
		responses = append(responses, models.FileUploadResponse{
			ID:           file.ID,
			FileName:     file.FileName,
			OriginalName: file.OriginalName,
			FileSize:     file.FileSize,
			MimeType:     file.MimeType,
			DownloadURL:  fmt.Sprintf("/v2/files/%d/download", file.ID),
			CreatedAt:    file.CreatedAt,
			Category:     file.Category,
			Description:  file.Description,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data":        responses,
		"page":        query.Page,
		"page_size":   query.PageSize,
		"total_items": total,
	})
}

func DeleteFile(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file ID"})
		return
	}

	var fileUpload models.FileUpload
	if err := core.DB.First(&fileUpload, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	userID := c.GetUint("user_id")
	if fileUpload.UploadedBy != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	if err := fileService.DeleteFile(fileUpload.FilePath); err != nil {
		core.ErrorLog("delete_file", err.Error())
	}

	if err := core.DB.Delete(&fileUpload).Error; err != nil {
		core.ErrorLog("delete_file", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete file record"})
		return
	}

	core.Log("delete_file", core.H{"file_id": id, "user_id": userID})
	c.JSON(http.StatusOK, gin.H{"message": "File deleted successfully"})
}
