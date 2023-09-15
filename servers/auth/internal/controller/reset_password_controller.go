package controller

import (
	database "auth-service/pkg/database"
	sqlc "auth-service/pkg/database/sqlc"
	utils "auth-service/pkg/utils"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type resetPasswordRequestBody struct {
	ResetToken  string `json:"reset_token" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,min=8,max=32,alphanum"`
}

type ResetPasswordController struct{}

func (controller *ResetPasswordController) updateNewPassword(
	c echo.Context,
	store database.Store,
	req resetPasswordRequestBody,
	userID string,
) (bool, error) {
	newPassword, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		log.Error(err)
		return false, utils.ErrInternalServerRepsonse()
	}

	err = store.ResetPassword(
		c.Request().Context(),
		sqlc.ResetPasswordParams{
			HashedPassword: newPassword,
			UserID:         userID,
		},
	)

	if err != nil {
		log.Error(err)
		return false, utils.ErrInternalServerRepsonse()
	}

	err = utils.RedisClient.Del(
		c.Request().Context(),
		req.ResetToken,
	).Err()

	if err != nil {
		log.Error(err)
		return false, utils.ErrInternalServerRepsonse()
	}

	return true, nil
}

func (controller *ResetPasswordController) checkResetToken(
	c echo.Context,
	req resetPasswordRequestBody,
) (string, bool, error) {
	userID, err := utils.RedisClient.Get(
		c.Request().Context(),
		req.ResetToken,
	).Result()

	if err != nil {
		log.Error(err)
		return "", false, utils.ErrInternalServerRepsonse()
	}

	return userID, true, nil
}

func (controller *ResetPasswordController) Execute(
	c echo.Context,
	store database.Store,
) error {
	var req resetPasswordRequestBody
	if err := c.Bind(&req); err != nil {
		log.Error(err)
		return err
	}
	if err := c.Validate(req); err != nil {
		log.Error(err)
		return err
	}

	userID, ok, err := controller.checkResetToken(c, req)
	if !ok {
		return err
	}

	ok, err = controller.updateNewPassword(c, store, req, userID)
	if !ok {
		return err
	}

	return c.JSON(200, utils.Response{
		Success: true,
		Message: "Reset password success",
		Payload: "",
	})
}
