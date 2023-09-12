package controller

import (
	"github.com/labstack/echo/v4"
)

type VerifyEmailController struct{}

func (controller VerifyEmailController) Execute(c echo.Context) error {
	return c.JSON(200, &map[string]interface{}{
		"success": true,
	})
}
