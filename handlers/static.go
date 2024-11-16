package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Home(c echo.Context) error {
	//#TODO: Exchange string returned on root with static about/intro page
	return c.Render(http.StatusOK, "index.html", nil)
}
