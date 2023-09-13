package controller

import (
	"github.com/labstack/echo/v4"
)

type resetPasswordRequest struct {
	NewPassword string `json:"new_password" validate:"required,min=8,max=32,alphanum"`
}

type ResetPasswordController struct{}

func (controller *ResetPasswordController) Validate(c echo.Context) (bool, error) {
	var req resetPasswordRequest
	if err := c.Bind(&req); err != nil {
		return false, err
	}
	if err := c.Validate(req); err != nil {
		return false, err
	}
	return true, nil
}

func (controller *ResetPasswordController) Execute(c echo.Context) error {
	if ok, err := controller.Validate(c); !ok {
		return err
	}

	return c.JSON(200, &Response{
		Success: true,
		Message: "Reset password success",
		Payload: "",
	})
}
