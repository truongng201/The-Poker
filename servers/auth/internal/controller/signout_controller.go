package controller

import (
	config "auth-service/config"
	database "auth-service/pkg/database"
	utils "auth-service/pkg/utils"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type signoutRequestBody struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
	AccessToken  string `json:"access_token" validate:"required"`
}

type SignoutController struct{}

func (controller *SignoutController) revokeAccessToken(
	c echo.Context,
	req signoutRequestBody,
) (string, bool, error) {
	claims, err := utils.GetJWTClaims(req.AccessToken, config.Con.JWT.SecretKey)
	if err != nil {
		log.Error(err)
		return "", false, utils.ErrTokenInvalidResponse()
	}
	accessTokenEmail := claims["sub"].(map[string]interface{})["email"].(string)
	return accessTokenEmail, true, nil
}

func (controller *SignoutController) revokeRefreshToken(
	c echo.Context,
	req signoutRequestBody,
	store database.Store,
	accessTokenEmail string,
) (bool, error) {
	email, err := utils.RedisClient.Get(c.Request().Context(), req.RefreshToken).Result()
	if err != nil {
		log.Error(err)
		return false, utils.ErrInternalServerResponse()
	}

	if email != accessTokenEmail {
		return false, utils.ErrWrongCredentialsResponse()
	}

	err = utils.RedisClient.Del(c.Request().Context(), req.RefreshToken).Err()
	if err != nil {
		log.Error(err)
		return false, utils.ErrInternalServerResponse()
	}

	err = store.DeleteRefreshToken(c.Request().Context(), req.RefreshToken)
	if err != nil {
		log.Error(err)
		return false, utils.ErrInternalServerResponse()
	}

	return true, nil
}

func (controller *SignoutController) Execute(
	c echo.Context,
	store database.Store,
) error {
	var req signoutRequestBody
	if err := c.Bind(&req); err != nil {
		log.Error(err)
		return err
	}
	if err := c.Validate(&req); err != nil {
		log.Error(err)
		return err
	}

	accessTokenEmail, ok, err := controller.revokeAccessToken(c, req)
	if !ok {
		return err
	}

	ok, err = controller.revokeRefreshToken(c, req, store, accessTokenEmail)
	if !ok {
		return err
	}

	return c.JSON(200, &utils.Response{
		Success: true,
		Message: "Sign out success",
		Payload: "",
	})
}
