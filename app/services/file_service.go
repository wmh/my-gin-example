package services

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	MaxFileSize       = 10 << 20 // 10 MB
	MaxBatchSize      = 50 << 20 // 50 MB total for multiple uploads
	UploadDir         = "./uploads"
	AllowedImageTypes = "image/jpeg,image/png,image/gif,image/webp"
	AllowedDocTypes   = "application/pdf,application/msword,application/vnd.openxmlformats-officedocument.wordprocessingml.document"
)

var (
	// DangerousExtensions - comprehensive list of potentially dangerous file extensions
	DangerousExtensions = []string{
		".exe", ".bat", ".cmd", ".sh", ".ps1", ".vbs", ".vbe",
		".jar", ".app", ".dmg", ".scr", ".msi", ".com", ".pif",
	}
	// AllowedExtensions - whitelist of safe file extensions
	AllowedExtensions = []string{
		".jpg", ".jpeg", ".png", ".gif", ".webp", ".bmp",
		".pdf", ".doc", ".docx", ".txt", ".csv", ".xls", ".xlsx",
		".zip", ".mp4", ".mp3", ".avi", ".mov",
	}
)

type FileService struct {
	uploadDir string
	maxSize   int64
}

func NewFileService() *FileService {
	return &FileService{
		uploadDir: UploadDir,
		maxSize:   MaxFileSize,
	}
}

func (fs *FileService) InitUploadDir() error {
	if _, err := os.Stat(fs.uploadDir); os.IsNotExist(err) {
		return os.MkdirAll(fs.uploadDir, 0755)
	}
	return nil
}

func (fs *FileService) ValidateFile(file *multipart.FileHeader, allowedTypes string) error {
	if file.Size > fs.maxSize {
		return fmt.Errorf("file size exceeds maximum allowed size of %d bytes", fs.maxSize)
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	ext = filepath.Clean(ext)

	// Check against dangerous extensions
	for _, dangerous := range DangerousExtensions {
		if ext == dangerous {
			return fmt.Errorf("file extension %s is not allowed", ext)
		}
	}

	// Whitelist validation
	isAllowed := false
	for _, allowed := range AllowedExtensions {
		if ext == allowed {
			isAllowed = true
			break
		}
	}
	if !isAllowed {
		return fmt.Errorf("file extension %s is not in the allowed list", ext)
	}

	// MIME type validation with content detection
	if allowedTypes != "" {
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		buffer := make([]byte, 512)
		_, err = src.Read(buffer)
		if err != nil && err != io.EOF {
			return err
		}

		detectedType := http.DetectContentType(buffer)
		declaredType := file.Header.Get("Content-Type")

		if !strings.Contains(allowedTypes, declaredType) || !strings.Contains(allowedTypes, detectedType) {
			return fmt.Errorf("file type mismatch: declared=%s, detected=%s", declaredType, detectedType)
		}
	}

	return nil
}

func (fs *FileService) GenerateUniqueFileName(originalName string) string {
	ext := filepath.Ext(originalName)
	ext = strings.ToLower(ext)
	ext = filepath.Clean(ext)
	
	// Remove any path separators from extension
	ext = strings.ReplaceAll(ext, string(filepath.Separator), "")
	ext = strings.ReplaceAll(ext, "/", "")
	ext = strings.ReplaceAll(ext, "\\", "")
	
	timestamp := time.Now().Unix()
	hash := sha256.Sum256([]byte(fmt.Sprintf("%s-%d", originalName, timestamp)))
	hashStr := hex.EncodeToString(hash[:])[:16]
	return fmt.Sprintf("%d-%s%s", timestamp, hashStr, ext)
}

func (fs *FileService) SaveFile(file *multipart.FileHeader) (string, string, error) {
	if err := fs.InitUploadDir(); err != nil {
		return "", "", err
	}

	fileName := fs.GenerateUniqueFileName(file.Filename)
	filePath := filepath.Join(fs.uploadDir, fileName)

	src, err := file.Open()
	if err != nil {
		return "", "", err
	}
	defer src.Close()

	dst, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return "", "", err
	}
	defer dst.Close()

	buf := make([]byte, 32*1024) // 32KB buffer
	if _, err = io.CopyBuffer(dst, src, buf); err != nil {
		return "", "", err
	}

	return fileName, filePath, nil
}

func (fs *FileService) DeleteFile(filePath string) error {
	absUploadDir, err := filepath.Abs(fs.uploadDir)
	if err != nil {
		return errors.New("failed to resolve upload directory")
	}
	absFilePath, err := filepath.Abs(filePath)
	if err != nil {
		return errors.New("failed to resolve file path")
	}
	
	// Ensure the file is within the upload directory
	if !strings.HasPrefix(absFilePath, absUploadDir+string(os.PathSeparator)) && absFilePath != absUploadDir {
		return errors.New("invalid file path: outside upload directory")
	}
	
	if _, err := os.Stat(absFilePath); os.IsNotExist(err) {
		return errors.New("file not found")
	}
	return os.Remove(absFilePath)
}

func (fs *FileService) GetFileInfo(filePath string) (os.FileInfo, error) {
	return os.Stat(filePath)
}
