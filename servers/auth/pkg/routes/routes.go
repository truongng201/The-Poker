package routes

import (
	controller "auth-service/pkg/controller"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo, controller controller.AppController) *echo.Echo {

	// Static Routes
	e.Static("/", "../../../client/build")
	// ApiRoutes
	api := e.Group("/auth")

	api.GET("/health", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	return e

}
