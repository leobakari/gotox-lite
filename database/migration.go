package database

import (
	"gorm.io/gorm"
	"log"
	"todo-demo/models" // Import the models package
)

func Migrate(db *gorm.DB) {
	// Run the migration for multiple models
	err := db.AutoMigrate(&models.Todo{})
	if err != nil {
		log.Fatalf("Error during migration: %v", err)
	}
}
