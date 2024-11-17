package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"todo-demo/database"
	"todo-demo/models"
)

func GetUsers(c echo.Context) error {
	var users []models.User

	if err := database.DB.Find(&users).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "ERROR: Fetching users failed",
		})
	}

	// Create Map with data
	data := map[string]interface{}{
		"users": users,
	}

	return c.Render(http.StatusOK, "userlist.html", data)
}
