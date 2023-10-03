package routes

import (
	controller "ws-service/internal/controller"
	socket "ws-service/pkg/socket"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo, controller *controller.AppController) *echo.Echo {
	api := e.Group("/ws")

	api.GET("/health", func(c echo.Context) error {
		return controller.HealthCheckController.Execute(c)
	})

	wsServer := socket.NewWsServer()
	go wsServer.Start()

	api.POST("/create-room", func(c echo.Context) error {
		return controller.CreateRoomController.Execute(c, wsServer)
	})

	api.POST("/join-room", func(c echo.Context) error {
		return controller.JoinRoomController.Execute(c, wsServer)
	})

	api.POST("/room", func(c echo.Context) error {
		return controller.WsRoomController.Execute(c, wsServer)
	})

	return e
}
