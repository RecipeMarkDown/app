package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/recipemarkdown/app/backend/internal/handlers"
)

// We pass in our db, from there we pass to the APIs
func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// allow requests from vite in local development (no reverse proxy)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))
	// group all of our route handlers under /api
	api := r.Group("/api")
	{
		// e.g. GET /api/test is handled by handlers.HelloHandler
		api.GET("/test", handlers.HelloHandler)

		// User routes
		api.POST("/users", handlers.CreateUser(db))
		api.GET("/users/:id", handlers.GetUserByID(db))
		api.GET("/users/email/:email", handlers.GetUserByEmail(db))
		api.GET("/users/google/:google_id", handlers.GetUserByGoogleID(db))
		api.PUT("/users/:id", handlers.UpdateUser(db))
	}

	return r
}
