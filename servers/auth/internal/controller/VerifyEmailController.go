package controller

import (
	"auth-service/pkg/utils"

	"github.com/labstack/echo/v4"
)

type verifyEmailRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type VerifyEmailController struct{}

func (controller *VerifyEmailController) Validate(c echo.Context) (bool, error) {
	var req verifyEmailRequest
	if err := c.Bind(&req); err != nil {
		return false, err
	}
	if err := c.Validate(&req); err != nil {
		return false, err
	}
	return true, nil
}

func (controller *VerifyEmailController) Execute(c echo.Context) error {
	if ok, err := controller.Validate(c); !ok {
		return err
	}

	return c.JSON(200, &utils.Response{
		Success: true,
		Message: "Verify email success",
		Payload: "",
	})

}
