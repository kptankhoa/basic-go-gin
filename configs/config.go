package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	PORT    string
	DB_HOST string
	DB_USER string
	DB_PASS string
)

// LoadEnv loads environment variables and initializes package variables
func LoadEnv(filenames ...string) {
	err := godotenv.Load(filenames...)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize package-level variables
	PORT = getEnv("PORT", "3000")
	DB_HOST = getEnv("DB_HOST", "127.0.0.1")
	DB_USER = getEnv("DB_USER", "root")
	DB_PASS = getEnv("DB_PASS", "password")
}

// getEnv fetches an environment variable with a fallback default
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
