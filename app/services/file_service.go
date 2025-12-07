package services

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	MaxFileSize       = 10 << 20 // 10 MB
	UploadDir         = "./uploads"
	AllowedImageTypes = "image/jpeg,image/png,image/gif,image/webp"
	AllowedDocTypes   = "application/pdf,application/msword,application/vnd.openxmlformats-officedocument.wordprocessingml.document"
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

	if allowedTypes != "" {
		mimeType := file.Header.Get("Content-Type")
		if !strings.Contains(allowedTypes, mimeType) {
			return fmt.Errorf("file type %s is not allowed", mimeType)
		}
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext == ".exe" || ext == ".sh" || ext == ".bat" {
		return errors.New("executable files are not allowed")
	}

	return nil
}

func (fs *FileService) GenerateUniqueFileName(originalName string) string {
	ext := filepath.Ext(originalName)
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

	dst, err := os.Create(filePath)
	if err != nil {
		return "", "", err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return "", "", err
	}

	return fileName, filePath, nil
}

func (fs *FileService) DeleteFile(filePath string) error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return errors.New("file not found")
	}
	return os.Remove(filePath)
}

func (fs *FileService) GetFileInfo(filePath string) (os.FileInfo, error) {
	return os.Stat(filePath)
}
