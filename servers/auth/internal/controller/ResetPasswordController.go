package controller

import (
	"github.com/labstack/echo/v4"
)

type ResetPasswordController struct{}

func (controller ResetPasswordController) Execute(c echo.Context) error {
	return c.JSON(200, &map[string]interface{}{
		"success": true,
	})
}
