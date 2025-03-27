package main

import (
	"log"

	"github.com/recipemarkdown/app/backend/internal/config"
	"github.com/recipemarkdown/app/backend/internal/routes"
)

func main() {
	// load environment variables from .env file if present
	config.LoadEnv()

	// get port from environment variable with fallback to 8000
	port := config.GetEnv("PORT", "8000")

	r := routes.SetupRouter()

	log.Printf("Starting server on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
