package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Greeting(c echo.Context) error {
	return c.String(http.StatusOK, "Hi visitor handler from the web side!")
}
