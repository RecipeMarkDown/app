package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/recipemarkdown/app/backend/internal/handlers"
)

func SetupRouter() *gin.Engine {
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
	}

	return r
}
