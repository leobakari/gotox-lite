package routes

import (
	"github.com/labstack/echo/v4"
	"todo-demo/handlers"
)

func RegisterRoutes(e *echo.Echo) {
	// Serving homepage
	e.GET("/", handlers.Home)
}
