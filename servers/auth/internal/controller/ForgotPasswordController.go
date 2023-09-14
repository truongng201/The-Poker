package controller

import (
	database "auth-service/pkg/database"
	utils "auth-service/pkg/utils"

	"github.com/labstack/echo/v4"
)

type forgotPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type ForgotPasswordController struct{}

func (controller *ForgotPasswordController) checkEmailExists(
	c echo.Context,
	store database.Store,
	req forgotPasswordRequest,
) (bool, error) {
	userInfo, err := store.FindUserByEmail(c.Request().Context(), req.Email)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return false, c.JSON(200, &utils.Response{
				Success: false,
				Message: "Email not found",
				Payload: "",
			})
		}
		return false, c.JSON(500, &utils.Response{
			Success: false,
			Message: "Internal server error",
			Payload: "",
		})
	}

	if !userInfo.IsVerified {
		return false, c.JSON(200, &utils.Response{
			Success: false,
			Message: "Email not verified",
			Payload: "",
		})
	}

	return true, nil
}

func (controller *ForgotPasswordController) Execute(
	c echo.Context,
	store database.Store,
) error {
	var req forgotPasswordRequest
	if err := c.Bind(&req); err != nil {
		return err
	}
	if err := c.Validate(&req); err != nil {
		return err
	}

	ok, err := controller.checkEmailExists(c, store, req)
	if !ok {
		return err
	}

	return c.JSON(200, &utils.Response{
		Success: true,
		Message: "Resend verification email success",
		Payload: "",
	})
}
