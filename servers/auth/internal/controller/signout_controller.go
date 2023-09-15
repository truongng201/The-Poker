package controller

import (
	"auth-service/pkg/utils"

	"github.com/labstack/echo/v4"
)

type signoutRequestBody struct {
	Email        string `json:"email" validate:"required,email"`
	RefreshToken string `json:"refresh_token" validate:"required"`
	AccessToken  string `json:"access_token" validate:"required"`
}

type SignoutController struct{}

func (controller *SignoutController) Validate(c echo.Context) (bool, error) {
	var req signoutRequestBody
	if err := c.Bind(&req); err != nil {
		return false, err
	}
	if err := c.Validate(&req); err != nil {
		return false, err
	}
	return true, nil
}

func (controller *SignoutController) Execute(c echo.Context) error {
	if ok, err := controller.Validate(c); !ok {
		return err
	}

	return c.JSON(200, &utils.Response{
		Success: true,
		Message: "Sign out success",
		Payload: "",
	})
}
