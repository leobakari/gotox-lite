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

	fmt.Println("Hello world!")

	title := c.FormValue("title")
	description := c.FormValue("description")

	todo := &models.Todo{
		Title:       title,
		Description: description,
	}

	fmt.Println(todo)

	if err := database.DB.Create(todo).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "ERROR: Creating Todo failed",
		})
	}
	return c.JSON(http.StatusCreated, todo)
}

func GetTodos(c echo.Context) error {
	var todos []models.Todo

	if err := database.DB.Find(&todos).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "ERROR: Fetching Todos failed",
		})
	}

	// Create Map with data
	data := map[string]interface{}{
		"todos": todos,
	}

	return c.Render(http.StatusOK, "todolist.html", data)
}
