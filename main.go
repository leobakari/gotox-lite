package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"todo-demo/db"
	"todo-demo/models"
	"todo-demo/routers"
)

func main() {

	// Creating a new instance of Echo
	e := echo.New()

	// Adding Template Parsing
	templates := models.NewTemplates("templates/*.html")
	e.Renderer = templates

	// Initiating Middleware
	// #TODO: Finding a visible more pleasing formating of logs
	e.Use(middleware.Logger())

	// SQLite database connection
	db, err := gorm.Open(sqlite.Open("todo-demo.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error opening database:", err)
	}

	// Run migrations for all models - Comes from /db
	migrations.Migrate(db)

	// Registering the Routes
	routes.RegisterRoutes(e)

	// Starting under port 5173 and logging fatal error.
	// 5173 used across all projects (coming from svelte)
	e.Logger.Fatal(e.Start(":5173"))

}
