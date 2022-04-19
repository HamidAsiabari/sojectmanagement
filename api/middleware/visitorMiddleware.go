package middleware

import (
	"github.com/labstack/echo/v4"
)

func SetVisitorMiddlewares(e *echo.Echo) {
	// e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
	// 	Root: "../static",
	// }))

	// e.Use(serverHeader)
}

func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "BlueBot/1.0")
		c.Response().Header().Set("notReallyHeader", "thisHaveNoMeaning")

		return next(c)
	}
}
