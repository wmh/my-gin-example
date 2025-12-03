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

func setupUserTestDB(t *testing.T) {
	var err error
	core.DB, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to initialize test database: %v", err)
	}
	if err := core.DB.AutoMigrate(&models.User{}); err != nil {
		t.Fatalf("Failed to migrate test database: %v", err)
	}
}

func TestRegister(t *testing.T) {
	setupUserTestDB(t)
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.POST("/register", Register)

	userReq := models.UserRegisterRequest{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
		FullName: "Test User",
	}

	jsonData, _ := json.Marshal(userReq)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "User registered successfully", response["message"])
}

func TestLogin(t *testing.T) {
	setupUserTestDB(t)
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.POST("/register", Register)
	r.POST("/login", Login)

	userReq := models.UserRegisterRequest{
		Username: "loginuser",
		Email:    "login@example.com",
		Password: "password123",
		FullName: "Login User",
	}

	jsonData, _ := json.Marshal(userReq)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	loginReq := models.UserLoginRequest{
		Username: "loginuser",
		Password: "password123",
	}

	jsonData, _ = json.Marshal(loginReq)
	req, _ = http.NewRequest("POST", "/login", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.NotEmpty(t, response["token"])
}
