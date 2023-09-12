package controller

import (
	"github.com/labstack/echo/v4"
)

type ForgotPasswordController struct{}

func (controller ForgotPasswordController) Execute(c echo.Context) error {
	return c.JSON(200, &map[string]interface{}{
		"success": true,
	})
}
