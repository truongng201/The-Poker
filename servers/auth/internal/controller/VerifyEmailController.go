package controller

import (
	database "auth-service/pkg/database"
	utils "auth-service/pkg/utils"

	"github.com/labstack/echo/v4"
)

type verifyEmailRequestParam struct {
	Token string `query:"token" validate:"required"`
}

type VerifyEmailController struct{}

func (controller *VerifyEmailController) checkToken(
	c echo.Context,
	req verifyEmailRequestParam,
) (string, bool, error) {
	email, err := utils.RedisClient.Get(c.Request().Context(), req.Token).Result()
	if err != nil {
		return "", false, c.JSON(200, &utils.Response{
			Success: false,
			Message: "Invalid token",
			Payload: "",
		})
	}
	return email, true, nil
}

func (controller *VerifyEmailController) updateIsVerified(
	c echo.Context,
	store database.Store,
	email string,
	req verifyEmailRequestParam,
) (bool, error) {
	err := store.VerifyEmail(c.Request().Context(), email)
	if err != nil {
		return false, c.JSON(500, &utils.Response{
			Success: false,
			Message: "Internal server error",
			Payload: "",
		})
	}
	err = utils.RedisClient.Del(c.Request().Context(), req.Token).Err()
	if err != nil {
		return false, c.JSON(500, &utils.Response{
			Success: false,
			Message: "Internal server error",
			Payload: "",
		})
	}
	return true, nil
}

func (controller *VerifyEmailController) Execute(c echo.Context, store database.Store) error {
	var reqParam verifyEmailRequestParam
	if err := c.Bind(&reqParam); err != nil {
		return err
	}
	if err := c.Validate(&reqParam); err != nil {
		return err
	}

	email, ok, err := controller.checkToken(c, reqParam)
	if !ok {
		return err
	}

	ok, err = controller.updateIsVerified(c, store, email, reqParam)
	if !ok {
		return err
	}

	return c.JSON(200, &utils.Response{
		Success: true,
		Message: "Verify email success",
		Payload: "",
	})

}
