package api

import (
	"github.com/HamidAsiabari/sojectmanagement/api/handler"

	"github.com/labstack/echo/v4"
)

func VisitorGroup(e *echo.Group) {
	e.GET("/", handler.Greeting)
	e.GET("/login", handler.Greeting)
	e.GET("/yallo", handler.Greeting)
	e.GET("/cats/:data", handler.Greeting)

	e.POST("/cats", handler.Greeting)
	e.POST("/dogs", handler.Greeting)
	e.POST("/hamsters", handler.Greeting)
}
