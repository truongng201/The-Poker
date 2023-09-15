package controller

import (
	config "auth-service/config"
	database "auth-service/pkg/database"
	utils "auth-service/pkg/utils"

	"fmt"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
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
		log.Error(err)
		return "", false, utils.ErrInternalServerRepsonse()
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
		log.Error(err)
		return false, utils.ErrInternalServerRepsonse()
	}
	err = utils.RedisClient.Del(c.Request().Context(), req.Token).Err()
	if err != nil {
		log.Error(err)
		return false, utils.ErrInternalServerRepsonse()
	}
	return true, nil
}

func (controller *VerifyEmailController) Execute(c echo.Context, store database.Store) error {
	var reqParam verifyEmailRequestParam
	if err := c.Bind(&reqParam); err != nil {
		log.Error(err)
		return utils.ErrBadRequestResponse()
	}
	if err := c.Validate(&reqParam); err != nil {
		log.Error(err)
		return utils.ErrBadRequestResponse()
	}

	email, ok, err := controller.checkToken(c, reqParam)
	if !ok {
		return err
	}

	ok, err = controller.updateIsVerified(c, store, email, reqParam)
	if !ok {
		return err
	}

	return c.Redirect(302, fmt.Sprintf("%s/signin", config.Con.Domains.Client))
}
