package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// environment variables already set will take precedence over .env values
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	} else {
		log.Println("Loaded environment from .env file")
	}
}

func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
