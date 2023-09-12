package controller

import (
	"github.com/labstack/echo/v4"
)

type SigninController struct{}

func (controller SigninController) Execute(c echo.Context) error {
	return c.JSON(200, &map[string]interface{}{
		"success": true,
	})
}
