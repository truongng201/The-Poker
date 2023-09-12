package controller

import (
	"github.com/labstack/echo/v4"
)

type SignupController struct{}

func (controller SignupController) Execute(c echo.Context) error {
	return c.JSON(200, &map[string]interface{}{
		"success": true,
	})
}
