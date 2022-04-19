package router

import (
	"github.com/HamidAsiabari/sojectmanagement/api"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	// create groups
	visitorGroup := e.Group("/visitor")
	// set all middlewares
	// middleware.SetVisitorMiddlewares(e)

	// set group routes
	api.VisitorGroup(visitorGroup)

	return e
}
