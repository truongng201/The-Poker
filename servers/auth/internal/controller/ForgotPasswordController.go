package controller

import (
	database "auth-service/pkg/database"
	sqlc "auth-service/pkg/database/sqlc"
	templates "auth-service/pkg/templates/email"
	utils "auth-service/pkg/utils"

	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type forgotPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type ForgotPasswordController struct{}

func (controller *ForgotPasswordController) checkEmailExists(
	c echo.Context,
	store database.Store,
	req forgotPasswordRequest,
) (sqlc.FindUserByEmailRow ,bool, error) {
	userInfo, err := store.FindUserByEmail(c.Request().Context(), req.Email)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return userInfo, false, c.JSON(200, &utils.Response{
				Success: false,
				Message: "Email not found",
				Payload: "",
			})
		}
		return userInfo, false, c.JSON(500, &utils.Response{
			Success: false,
			Message: "Internal server error",
			Payload: "",
		})
	}

	if !userInfo.IsVerified {
		return userInfo, false, c.JSON(200, &utils.Response{
			Success: false,
			Message: "Email not verified",
			Payload: "",
		})
	}

	return userInfo, true, nil
}

func (controller *ForgotPasswordController) sendResetPasswordEmail(
	c echo.Context,
	store database.Store,
	req forgotPasswordRequest,
	mailer utils.EmailSender,
	userInfo sqlc.FindUserByEmailRow,
) (bool, error) {
	// Generate token
	token := uuid.New().String()

	// Save token to cache
	err := utils.RedisClient.Set(
		c.Request().Context(),
		token, 
		userInfo.UserID, 
		time.Duration(15) * time.Minute,
	).Err()
	if err != nil {
		log.Error(err)
		return false, c.JSON(500, &utils.Response{
			Success: false,
			Message: "Internal server error",
			Payload: "",
		})
	}

	// Send email
	err = mailer.SendEmail(
		"Reset password",
		templates.GenerateResetPasswordTemplate(
			templates.ResetPasswordTemplateData{
				Username: userInfo.Username,
				ResetLink: fmt.Sprintf("https://beta.truongng.me/reset/%s", token),
			},
		),
		[]string{userInfo.Email},
		[]string{},
		[]string{},
		[]string{},
	)
	if err != nil {
		log.Error(err)
		return false, c.JSON(500, &utils.Response{
			Success: false,
			Message: "Internal server error",
			Payload: "",
		})
	}

	return true, nil
}

func (controller *ForgotPasswordController) Execute(
	c echo.Context,
	store database.Store,
	mailer utils.EmailSender,
) error {
	var req forgotPasswordRequest
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

	ok, err = controller.sendResetPasswordEmail(c, store, req, mailer, userInfo)
	if !ok {
		return err
	}

	return c.JSON(200, &utils.Response{
		Success: true,
		Message: "Send reset password email success",
		Payload: "",
	})
}
