package database

import (
	"errors"
	"fmt"
	"kptankhoa.dev/basic-go-gin/configs"
	"kptankhoa.dev/basic-go-gin/internal/auth"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

var errCannotConnectToDB = errors.New("cannot connect to database")

func ConnectDatabase() error {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s search_path=%s sslmode=disable TimeZone=UTC",
		configs.DB_HOST,
		configs.DB_USER,
		configs.DB_PASS,
		configs.DB_NAME,
		configs.DB_PORT,
		configs.DB_SCHEMA,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Enable logging
	})

	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
		return errCannotConnectToDB
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get SQL DB:", err)
		return errCannotConnectToDB
	}

	// Set connection pool settings
	sqlDB.SetMaxOpenConns(25)                 // Maximum open connections
	sqlDB.SetMaxIdleConns(10)                 // Maximum idle connections
	sqlDB.SetConnMaxLifetime(5 * time.Minute) // Connection lifetime

	DB = db
	fmt.Println("Database connected successfully")

	autoMigrate(db)

	return nil
}

func autoMigrate(db *gorm.DB) {
	db.AutoMigrate(&auth.User{})
}
