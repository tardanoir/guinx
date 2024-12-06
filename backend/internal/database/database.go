package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"guinx/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	// Get database URL from environment variable
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return nil, fmt.Errorf("DATABASE_URL environment variable is not set")
	}

	// Try to connect with retries
	var db *gorm.DB
	var err error
	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("Failed to connect to database (attempt %d/%d): %v", i+1, maxRetries, err)
		time.Sleep(time.Second * 5) // Wait 5 seconds before retrying
	}
	if err != nil {
		return nil, err
	}

	// Auto migrate the schema
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
