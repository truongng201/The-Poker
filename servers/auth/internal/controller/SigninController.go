package controller

import (
	"github.com/labstack/echo/v4"
)

type signinRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32,alphanum"`
}

type signinResponsePayload struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

type SigninController struct{}

func (controller *SigninController) Validate(c echo.Context) (bool, error) {
	var req signinRequest
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

func (controller *SigninController) Execute(c echo.Context) error {
	if ok, err := controller.Validate(c); !ok {
		return err
	}

	return c.JSON(200, &Response{
		Success: true,
		Message: "Sign in success",
		Payload: &signinResponsePayload{
			AccessToken:  "access_token",
			RefreshToken: "refresh_token",
		},
	})
}
