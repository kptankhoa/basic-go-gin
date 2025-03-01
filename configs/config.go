package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	PORT       string
	DB_HOST    string
	DB_PORT    string
	DB_USER    string
	DB_PASS    string
	DB_NAME    string
	DB_SCHEMA  string
	JWT_SECRET string
)

// LoadEnv loads environment variables and initializes package variables
func LoadEnv(filenames ...string) {
	err := godotenv.Load(filenames...)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize package-level variables
	PORT = getEnv("PORT", "8080")
	DB_HOST = getEnv("DB_HOST", "127.0.0.1")
	DB_PORT = getEnv("DB_PORT", "5432")
	DB_USER = getEnv("DB_USER", "postgres")
	DB_PASS = getEnv("DB_PASS", "postgres")
	DB_NAME = getEnv("DB_NAME", "music")
	DB_SCHEMA = getEnv("DB_SCHEMA", "")
	JWT_SECRET = getEnv("JWT_SECRET", "")
}

// getEnv fetches an environment variable with a fallback default
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
