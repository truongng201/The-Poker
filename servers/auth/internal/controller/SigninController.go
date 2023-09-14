package controller

import (
	database "auth-service/pkg/database/sqlc"
	"auth-service/pkg/utils"

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

func (controller *SigninController) checkEmailExists(
	c echo.Context,
	store database.Store,
	req signinRequest,
) (database.GetUserByEmailRow, bool, error) {
	res, err := store.GetUserByEmail(c.Request().Context(), req.Email)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return res, false, c.JSON(200, &utils.Response{
				Success: false,
				Message: "Email not found",
				Payload: "",
			})
		}
		return res, false, c.JSON(500, &utils.Response{
			Success: false,
			Message: "Internal server error",
			Payload: "",
		})
	}

	return res, true, nil
}

func (controller *SigninController) checkPassword(
	c echo.Context, 
	store database.Store, 
	req signinRequest, 
	res database.GetUserByEmailRow,
) (bool, error) {
	if !utils.CheckPassword(req.Password, res.HashedPassword) {
		return false, c.JSON(200, &utils.Response{
			Success: false,
			Message: "Invalid password",
			Payload: "",
		})
	}

	return true, nil
}

func (controller *SigninController) Execute(c echo.Context, store database.Store) error {
	var req signinRequest

	if err := c.Bind(&req); err != nil {
		return err
	}
	if err := c.Validate(&req); err != nil {
		return err
	}

	res, ok, err := controller.checkEmailExists(c, store, req)
	if !ok {
		return err
	}

	ok, err = controller.checkPassword(c, store, req, res)
	if !ok {
		return err
	}

	return c.JSON(200, &utils.Response{
		Success: true,
		Message: "Sign in success",
		Payload: &signinResponsePayload{
			AccessToken:  "access_token",
			RefreshToken: "refresh_token",
		},
	})
}
