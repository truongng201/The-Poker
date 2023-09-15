package controller

import (
	"fmt"
	"time"

	config "auth-service/config"
	database "auth-service/pkg/database"
	sqlc "auth-service/pkg/database/sqlc"
	templates "auth-service/pkg/templates/email"
	utils "auth-service/pkg/utils"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type forgotPasswordRequestBody struct {
	Email string `json:"email" validate:"required,email"`
}

type ForgotPasswordController struct{}

func (controller *ForgotPasswordController) checkEmailExists(
	c echo.Context,
	store database.Store,
	req forgotPasswordRequestBody,
) (sqlc.GetUserByEmailRow, bool, error) {
	userInfo, err := store.GetUserByEmail(c.Request().Context(), req.Email)

	if err != nil {
		if err.Error() == "no rows in result set" {
			return userInfo, false, utils.ErrNoSuchUserResponse()
		}
		log.Error(err)
		return userInfo, false, utils.ErrInternalServerRepsonse()
	}

	if !userInfo.IsVerified {
		return userInfo, false, c.Redirect(302, fmt.Sprintf("%s/reverify", config.Con.Domains.Client))
	}

	return userInfo, true, nil
}

func (controller *ForgotPasswordController) sendResetPasswordEmail(
	c echo.Context,
	store database.Store,
	req forgotPasswordRequestBody,
	mailer utils.EmailSender,
	userInfo sqlc.GetUserByEmailRow,
	resetPasswordtoken string,
) (bool, error) {
	err := mailer.SendEmail(
		"Reset password",
		templates.GenerateResetPasswordTemplate(
			templates.ResetPasswordTemplateData{
				Username:   userInfo.Username,
				ResetLink:  fmt.Sprintf("%s/reset/%s", config.Con.Domains.Client, resetPasswordtoken),
				ExpireTime: config.Con.Timeout.ResetPasswordToken,
			},
		),
		[]string{req.Email},
		[]string{},
		[]string{},
		[]string{},
	)

	if err != nil {
		log.Error(err)
		return false, utils.ErrBadRequestResponse()
	}

	return true, nil
}

func (controller *ForgotPasswordController) generateResetPasswordToken(
	c echo.Context,
	userInfo sqlc.GetUserByEmailRow,
) (string, bool, error) {
	resetPasswordToken := uuid.New().String()

	err := utils.RedisClient.Set(
		c.Request().Context(),
		resetPasswordToken,
		userInfo.UserID,
		time.Duration(config.Con.Timeout.ResetPasswordToken)*time.Minute,
	).Err()
	if err != nil {
		log.Error(err)
		return "", false, utils.ErrInternalServerResponse()
	}

	return resetPasswordToken, true, nil
}

func (controller *ForgotPasswordController) Execute(
	c echo.Context,
	store database.Store,
	mailer utils.EmailSender,
) error {
	var req forgotPasswordRequestBody
	if err := c.Bind(&req); err != nil {
		log.Error(err)
		return err
	}
	if err := c.Validate(&req); err != nil {
		return err
	}

	userInfo, ok, err := controller.checkEmailExists(c, store, req)
	if !ok {
		return err
	}

	resetPasswordToken, ok, err := controller.generateResetPasswordToken(c, userInfo)
	if !ok {
		return err
	}

	ok, err = controller.sendResetPasswordEmail(
		c,
		store,
		req,
		mailer,
		userInfo,
		resetPasswordToken,
	)
	if !ok {
		return err
	}

	return c.JSON(200, utils.Response{
		Success: true,
		Message: "Reset password email sent",
		Payload: "",
	})
}
