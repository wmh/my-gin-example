package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wmh/my-gin-example/app/core"
	"github.com/wmh/my-gin-example/app/models"
	"github.com/wmh/my-gin-example/app/services"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var req models.UserRegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.User
	if err := core.DB.Where("username = ? OR email = ?", req.Username, req.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Username or email already exists"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		core.ErrorLog("register", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
		FullName: req.FullName,
		Role:     "user",
		IsActive: true,
	}

	if err := core.DB.Create(&user).Error; err != nil {
		core.ErrorLog("register", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	core.Log("register", core.H{"user_id": user.ID, "username": user.Username})
	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user":    user.ToResponse(),
	})
}

func Login(c *gin.Context) {
	var req models.UserLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := core.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if !user.IsActive {
		c.JSON(http.StatusForbidden, gin.H{"error": "Account is disabled"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := services.GenerateToken(user.ID, user.Username, user.Role)
	if err != nil {
		core.ErrorLog("login", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	core.Log("login", core.H{"user_id": user.ID, "username": user.Username})
	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  user.ToResponse(),
	})
}

func GetProfile(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var user models.User
	if err := core.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user.ToResponse())
}

func UpdateProfile(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var req models.UserUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := core.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	updates := map[string]interface{}{}
	if req.Email != "" {
		updates["email"] = req.Email
	}
	if req.FullName != "" {
		updates["full_name"] = req.FullName
	}

	if len(updates) > 0 {
		if err := core.DB.Model(&user).Updates(updates).Error; err != nil {
			core.ErrorLog("update_profile", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
			return
		}
	}

	core.Log("update_profile", core.H{"user_id": user.ID})
	c.JSON(http.StatusOK, gin.H{
		"message": "Profile updated successfully",
		"user":    user.ToResponse(),
	})
}

func ListUsers(c *gin.Context) {
	var users []models.User
	if err := core.DB.Find(&users).Error; err != nil {
		core.ErrorLog("list_users", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	var responses []models.UserResponse
	for _, user := range users {
		responses = append(responses, user.ToResponse())
	}

	c.JSON(http.StatusOK, gin.H{
		"users": responses,
		"count": len(responses),
	})
}
