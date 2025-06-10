// internal/handlers/user_handler.go
package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/recipemarkdown/app/backend/internal/models"
	"gorm.io/gorm"
)

// Create - POST /users
func CreateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		c.ShouldBindJSON(&user)
		db.Create(&user)
		c.JSON(201, user)
	}
}

// GetByID - GET /users/:id
func GetUserByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var user models.User
		if err := db.First(&user, id).Error; err != nil {
			c.JSON(404, gin.H{"error": "User not found"})
			return
		}

		c.JSON(200, user)
	}
}

// GetByEmail - GET /users/email/:email
func GetUserByEmail(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.Param("email")

		var user models.User
		if err := db.Where("email = ?", email).First(&user).Error; err != nil {
			c.JSON(404, gin.H{"error": "User not found"})
			return
		}

		c.JSON(200, user)
	}
}

// GetByGoogleID - GET /users/google/:google_id
func GetUserByGoogleID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		googleID := c.Param("google_id")

		var user models.User
		if err := db.Where("google_id = ?", googleID).First(&user).Error; err != nil {
			c.JSON(404, gin.H{"error": "User not found"})
			return
		}

		c.JSON(200, user)
	}
}

// Update - PUT /users/:id
func UpdateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		// First, check if user exists
		var user models.User
		if err := db.First(&user, id).Error; err != nil {
			c.JSON(404, gin.H{"error": "User not found"})
			return
		}

		// Bind new data to user
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// Save updated user
		if err := db.Save(&user).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to update user"})
			return
		}

		c.JSON(200, user)
	}
}
