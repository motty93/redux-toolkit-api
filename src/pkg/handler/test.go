package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func TestHandler(c echo.Context) error {
	response := map[string]string{"ping": "pong"}

	return c.JSON(http.StatusOK, response)
}
