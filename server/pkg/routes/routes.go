package routes

import (
	controller "server/pkg/controller"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo, controller controller.AppController) *echo.Echo {
	// Routes
	api := e.Group("/api")

	api.GET("/health", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	return e

}