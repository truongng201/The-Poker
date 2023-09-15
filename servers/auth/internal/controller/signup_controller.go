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
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type signupRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32,alphanum"`
	Username string `json:"username" validate:"required,min=6,max=32"`
}

type signupResponsePayload struct {
	UserID string `json:"user_id"`
}

type SignupController struct{}

func (controller *SignupController) checkEmailExists(
	c echo.Context,
	store database.Store,
	req signupRequest,
) (bool, error) {
	_, err := store.GetUserByEmail(c.Request().Context(), req.Email)

	if err != nil {
		if err.Error() == "no rows in result set" {
			return true, nil
		}
		log.Error(err)
		return false, utils.ErrInternalServerRepsonse()
	}

	return false, utils.ErrEmailAlreadyExistsResponse()
}

func (controller *SignupController) createNewUser(
	c echo.Context,
	store database.Store,
	req signupRequest,
	hashedPassword string,
) (sqlc.CreateUserRow, bool, error) {
	userInfo, err := store.CreateUser(c.Request().Context(),
		sqlc.CreateUserParams{
			UserID:         uuid.New().String(),
			Username:       req.Username,
			Email:          req.Email,
			HashedPassword: hashedPassword,
			ImageUrl: pgtype.Text{
				String: fmt.Sprintf("https://api.dicebear.com/6.x/bottts-neutral/svg?seed=%s", req.Username),
				Valid:  true,
			},
		},
	)

	if err != nil {
		log.Error(err)
		return userInfo, false, utils.ErrInternalServerRepsonse()
	}

	return userInfo, true, nil
}

func (controller *SignupController) generateVerifyEmailToken(
	c echo.Context,
	userInfo sqlc.CreateUserRow,
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

func (controller *SignupController) sendVerifyEmail(
	c echo.Context,
	mailer utils.EmailSender,
	req signupRequest,
	userInfo sqlc.CreateUserRow,
	verifyToken string,
) (bool, error) {
	err := mailer.SendEmail(
		"Verify your email",
		templates.GenerateVerifyEmailTemplate(templates.VerifyEmailTemplateData{
			Username:   req.Username,
			VerifyLink: fmt.Sprintf("%s/verify-email?token=%s", config.Con.Domains.Server, verifyToken),
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

func (controller *SignupController) Execute(
	c echo.Context,
	store database.Store,
	mailer utils.EmailSender,
) error {
	var req signupRequest
	if err := c.Bind(&req); err != nil {
		log.Error(err)
		return err
	}
	if err := c.Validate(&req); err != nil {
		log.Error(err)
		return err
	}

	ok, err := controller.checkEmailExists(c, store, req)
	if !ok {
		return err
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return utils.ErrBadRequestResponse()
	}

	userInfo, ok, err := controller.createNewUser(c, store, req, hashedPassword)
	if !ok {
		return err
	}

	verifyToken, ok, err := controller.generateVerifyEmailToken(c, userInfo)
	if !ok {
		return err
	}

	ok, err = controller.sendVerifyEmail(c, mailer, req, userInfo, verifyToken)
	if !ok {
		return err
	}

	return utils.SuccessResponse(
		"Sign up success",
		&signupResponsePayload{
			UserID: userInfo.UserID,
		},
	)
}
