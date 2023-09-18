package controller

import "github.com/labstack/echo/v4"

type CreateRoomController struct{}

func (controller *CreateRoomController) Execute(c echo.Context) error {

	return c.JSON(200, &HealthCheck{
		Success: true,
		Message: "Create Room",
		Version: Version,
	})
}
