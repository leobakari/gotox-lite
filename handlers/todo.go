package handlers

import (
	"fmt"
	"net/http"
	"todo-demo/database"
	"todo-demo/models"

	"github.com/labstack/echo/v4"
)

func CreateTodo(c echo.Context) error {

	// Binding with c.Bind() didnt work, manual binding as work around
	//#TODO: Find a way to make automatic binding work (most likely because of the request content type)

	title := c.FormValue("title")
	description := c.FormValue("description")
	isActive := true

	todo := &models.Todo{
		Title:       title,
		Description: description,
		IsActive:    isActive,
	}

	if err := database.DB.Create(todo).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "ERROR: Creating Todo failed",
		})
	}
	// I had to change to the template after outsourcing the Todo
	// from the todolist in its own file - #NOTE: do not change
	return c.Render(http.StatusCreated, "Todo", todo)
}

func CloseTodo(c echo.Context) error {
	// NOTE: This function only soft deletes the todos / Deactivates them!
	// Grab ID from Request
	id := c.Param("id")
	fmt.Println(id)

	if err := database.DB.Model(&models.Todo{}).Where("id = ?", id).Update("is_active", 0).Error; err != nil {
		fmt.Println("Error deactivating todo:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to deactivate todo",
		})
	}

	return c.NoContent(http.StatusOK)
}

func GetTodos(c echo.Context) error {
	var todos []models.Todo

	if err := database.DB.Where("is_active", 1).Find(&todos).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "ERROR: Fetching Todos failed",
		})
	}

	// Format CreatedAt for each todo
	formattedTodos := make([]map[string]interface{}, len(todos))
	for i, todo := range todos {
		formattedTodos[i] = map[string]interface{}{
			"ID":          todo.ID,
			"Title":       todo.Title,
			"Description": todo.Description,
			"CreatedAt":   todo.CreatedAt.Format("2006-01-02 15:04"), // Format the date
		}
	}

	// Pass formatted data to the template
	data := map[string]interface{}{
		"todos": formattedTodos,
	}

	return c.Render(http.StatusOK, "todolist.html", data)
}
