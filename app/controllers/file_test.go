package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/wmh/my-gin-example/app/core"
	"github.com/wmh/my-gin-example/app/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupFileTestDB(t *testing.T) {
	var err error
	core.DB, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to init test DB: %v", err)
	}
	if err := core.DB.AutoMigrate(&models.User{}, &models.FileUpload{}); err != nil {
		t.Fatalf("Failed to migrate: %v", err)
	}
}

func TestUploadFile(t *testing.T) {
	t.Skip("Skipping file upload test - requires file system setup")
}

func TestListFiles(t *testing.T) {
	setupFileTestDB(t)
	defer core.CloseDB()

	hashedPassword2, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	user := models.User{
		Username: "testuser",
		Email:    "test@example.com",
		Password: string(hashedPassword2),
	}
	core.DB.Create(&user)

	fileUpload := models.FileUpload{
		FileName:     "test.txt",
		OriginalName: "original.txt",
		FilePath:     "/tmp/test.txt",
		FileSize:     1024,
		MimeType:     "text/plain",
		UploadedBy:   user.ID,
	}
	core.DB.Create(&fileUpload)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = httptest.NewRequest("GET", "/v2/files?page=1&page_size=10", nil)
	c.Set("user_id", user.ID)

	ListFiles(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "test.txt")
}
