// config/database.go
package config

import (
	"fmt"
	"log"

	"github.com/recipemarkdown/app/backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connects the Database
func ConnectDatabase() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		GetEnv("DB_HOST", "localhost"),
		GetEnv("DB_USER", "user"),
		GetEnv("DB_PASSWORD", "password"),
		GetEnv("DB_NAME", "recipemd"),
		GetEnv("DB_PORT", "5432"),
		GetEnv("DB_SSLMODE", "disable"),
		GetEnv("DB_TIMEZONE", "UTC"),
	)

	// Attempts to connect GORM to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	/*
		Attempts to AutoMigrate
		With each struct/database table that we have defined in /models, put them here.
		https://gorm.io/docs/migration.html
	*/
	if err := db.AutoMigrate(&models.User{}); err != nil {
		return nil, fmt.Errorf("migration failed: %w", err)
	}

	log.Println("Database connected successfully, gogogogogo")
	return db, nil
}
