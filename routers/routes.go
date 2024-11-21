package routes

import (
	"github.com/labstack/echo/v4"
	"todo-demo/handlers"
)

func RegisterRoutes(e *echo.Echo) {
	// Serving homepage
	e.GET("/", handlers.Home)

	// Todo related routes
	e.POST("/todo", handlers.CreateTodo)
	e.DELETE("/todo/:id", handlers.CloseTodo)

	e.GET("/todos", handlers.GetTodos)
}
