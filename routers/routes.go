package routes

import (
	"github.com/labstack/echo/v4"
	"todo-demo/handlers"
)

func RegisterRoutes(e *echo.Echo) {
	// Serving homepage
	e.GET("/", handlers.Home)
	e.GET("/click", handlers.Click)

	// User related routes
	e.GET("/users", handlers.GetUsers)

	// Todo related routes
	e.POST("/todo", handlers.CreateTodo)
	e.DELETE("/todo/:id", handlers.CloseTodo)

	e.GET("/todos", handlers.GetTodos)
}
