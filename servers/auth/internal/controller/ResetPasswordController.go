package controller

import (
	"auth-service/pkg/utils"

	"github.com/labstack/echo/v4"
)

type resetPasswordRequest struct {
	NewPassword string `json:"new_password" validate:"required,min=8,max=32,alphanum"`
}

type ResetPasswordController struct{}

func (controller *ResetPasswordController) Execute(c echo.Context) error {
	var req resetPasswordRequest
	if err := c.Bind(&req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}

	return c.JSON(200, &utils.Response{
		Success: true,
		Message: "Reset password success",
		Payload: "",
	})
}
