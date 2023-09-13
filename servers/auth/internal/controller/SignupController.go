package controller

import (
	"github.com/labstack/echo/v4"
)

type signupRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32,alphanum"`
	Username string `json:"username" validate:"required,min=6,max=32"`
}

type SignupController struct{}

func (controller *SignupController) Validate(c echo.Context) (bool, error) {
	var req signupRequest
	if err := c.Bind(&req); err != nil {
		return false, c.JSON(400, &Response{
			Success: false,
			Message: "Invalid request body",
			Payload: "",
		})
	}
	if err := c.Validate(req); err != nil {
		return false, c.JSON(400, &Response{
			Success: false,
			Message: "Invalid request body",
			Payload: "",
		})
	}
	return true, nil
}

func (controller SignupController) Execute(c echo.Context) error {
	if ok, err := controller.Validate(c); !ok {
		return err
	}

	return c.JSON(200, &Response{
		Success: true,
		Message: "Sign up success",
		Payload: "",
	})
}
