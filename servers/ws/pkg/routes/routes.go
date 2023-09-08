package routes

import (
	controller "ws-service/pkg/controller"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo, controller controller.AppController) *echo.Echo {
	api := e.Group("/ws")

	api.GET("/health", func(c echo.Context) error {
		return c.String(200, "Service is up and running!")
	})

	return e
}