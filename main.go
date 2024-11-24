package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"todo-demo/database"
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

	// Initiate Database
	database.InitiateDatabase()

	// Registering the Routes
	routes.RegisterRoutes(e)

	// Change here for localhost for local usage
	e.Logger.Fatal(e.Start("0.0.0.0:8000"))

}
