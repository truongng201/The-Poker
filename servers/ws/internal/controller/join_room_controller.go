package controller

import "github.com/labstack/echo/v4"

type JoinRoomController struct{}

func (controller *JoinRoomController) Execute(c echo.Context) error {
	return c.JSON(200, &HealthCheck{
		Success: true,
		Message: "Join Room",
		Version: Version,
	})
}
