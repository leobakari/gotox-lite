package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Click(c echo.Context) error {
	return c.String(http.StatusOK, "greets from the server")
}
