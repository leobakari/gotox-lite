package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitiateDatabase() {

	// SQLite database connection
	db, err := gorm.Open(sqlite.Open("todo-demo.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error opening database:", err)
	}

	DB = db

	// Run migrations for all models - Comes from /db
	Migrate(db)
}
