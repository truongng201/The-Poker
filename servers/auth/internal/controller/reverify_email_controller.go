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

type reverifyEmailRequestBody struct {
	Email string `json:"email" validate:"required,email"`
}

type ReverifyEmailController struct{}

func (controller *ReverifyEmailController) checkEmailExists(
	c echo.Context,
	store database.Store,
	req reverifyEmailRequestBody,
) (sqlc.GetUserByEmailRow, bool, error) {
	userInfo, err := store.GetUserByEmail(c.Request().Context(), req.Email)

	if err != nil {
		if err.Error() == "no rows in result set" {
			return userInfo, false, utils.ErrNoSuchUserResponse()
		}
		log.Error(err)
		return userInfo, false, utils.ErrInternalServerRepsonse()
	}

	if userInfo.IsVerified {
		return userInfo, false, utils.ErrBadRequestResponse()
	}

	return userInfo, true, nil
}

func (controller *ReverifyEmailController) generateVerifyEmailToken(
	c echo.Context,
	userInfo sqlc.GetUserByEmailRow,
) (string, bool, error) {
	verifyEmailToken := uuid.New().String()
	err := utils.RedisClient.Set(
		c.Request().Context(),
		verifyEmailToken,
		userInfo.Email,
		time.Duration(config.Con.Timeout.VerifyEmailToken)*time.Minute,
	).Err()

	if err != nil {
		log.Error(err)
		return "", false, utils.ErrInternalServerResponse()
	}

	return verifyEmailToken, true, nil
}

func (controller *ReverifyEmailController) sendVerificationEmail(
	c echo.Context,
	store database.Store,
	req reverifyEmailRequestBody,
	mailer utils.EmailSender,
	userInfo sqlc.GetUserByEmailRow,
	verifyEmailToken string,
) (bool, error) {
	err := mailer.SendEmail(
		"Verify your email",
		templates.GenerateVerifyEmailTemplate(templates.VerifyEmailTemplateData{
			Username:   userInfo.Username,
			VerifyLink: fmt.Sprintf("%s/auth/verify-email?token=%s", config.Con.Domains.Server, verifyEmailToken),
			ExpireTime: config.Con.Timeout.VerifyEmailToken,
		}),
		[]string{userInfo.Email},
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

func (controller *ReverifyEmailController) Execute(
	c echo.Context,
	store database.Store,
	mailer utils.EmailSender,
) error {
	var req reverifyEmailRequestBody
	if err := c.Bind(&req); err != nil {
		log.Error(err)
		return err
	}
	if err := c.Validate(&req); err != nil {
		log.Error(err)
		return err
	}

	userInfo, ok, err := controller.checkEmailExists(c, store, req)
	if !ok {
		return err
	}

	verifyEmailToken, ok, err := controller.generateVerifyEmailToken(c, userInfo)
	if !ok {
		return err
	}

	ok, err = controller.sendVerificationEmail(c, store, req, mailer, userInfo, verifyEmailToken)
	if !ok {
		return err
	}

	return c.JSON(200, utils.Response{
		Success: true,
		Message: "Verification email resent",
		Payload: "",
	})
}
