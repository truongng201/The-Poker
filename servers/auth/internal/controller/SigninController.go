package controller

import (
	database "auth-service/pkg/database"
	sqlc "auth-service/pkg/database/sqlc"
	"auth-service/pkg/utils"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
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
) (sqlc.GetUserByEmailRow, bool, error) {
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
	res sqlc.GetUserByEmailRow,
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

func (controller *SigninController) generateNewRefreshToken(
	c echo.Context,
	store database.Store,
	userInfo sqlc.GetUserByEmailRow,
) (sqlc.CreateRefreshTokenRow, bool, error) {
	res, err := store.CreateRefreshToken(c.Request().Context(), sqlc.CreateRefreshTokenParams{
		UserID: userInfo.ID,
		DeviceType: pgtype.Text{
			String: "device_type",
			Valid:  true,
		},
		UserAgent: pgtype.Text{
			String: "user_agent",
			Valid:  true,
		},
		IpAddress: "ip_address",
		Token:     "token",
	})
	if err != nil {
		return res, false, c.JSON(200, &utils.Response{
			Success: false,
			Message: "Sign in failed",
			Payload: "",
		})
	}

	return res, true, nil
}

func (controller *SigninController) Execute(c echo.Context, store database.Store) error {
	var req signinRequest

	if err := c.Bind(&req); err != nil {
		return err
	}
	if err := c.Validate(&req); err != nil {
		return err
	}

	userInfo, ok, err := controller.checkEmailExists(c, store, req)
	if !ok {
		return err
	}

	ok, err = controller.checkPassword(c, store, req, userInfo)
	if !ok {
		return err
	}

	tokenInfo, ok, err := controller.generateNewRefreshToken(c, store, userInfo)
	if !ok {
		return err
	}
	log.Info(tokenInfo)

	return c.JSON(200, &utils.Response{
		Success: true,
		Message: "Sign in success",
		Payload: &signinResponsePayload{
			AccessToken:  "access_token",
			RefreshToken: "refresh_token",
		},
	})
}
