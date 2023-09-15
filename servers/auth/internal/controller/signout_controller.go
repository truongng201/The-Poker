package controller

import (
	config "auth-service/config"
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
) (bool, error) {
	claims, err := utils.GetJWTClaims(req.AccessToken, config.Con.JWT.SecretKey)
	if err != nil {
		log.Error(err)
		return false, utils.ErrInternalServerResponse()
	}
	log.Info(claims)
	return true, nil
}

func (controller *SignoutController) Execute(c echo.Context) error {
	var req signoutRequestBody
	if err := c.Bind(&req); err != nil {
		log.Error(err)
		return err
	}
	if err := c.Validate(&req); err != nil {
		log.Error(err)
		return err
	}

	ok, err := controller.revokeAccessToken(c, req)
	if !ok {
		return err
	}

	return c.JSON(200, &utils.Response{
		Success: true,
		Message: "Sign out success",
		Payload: "",
	})
}
