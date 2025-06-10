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

	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	r := routes.SetupRouter(db)

	log.Println("Database setup and ready")

	log.Printf("Starting server on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
