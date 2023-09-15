package controller

import (
	database "auth-service/pkg/database"
	sqlc "auth-service/pkg/database/sqlc"
	"auth-service/pkg/utils"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
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
	_, err := store.CheckEmailExists(c.Request().Context(), req.Email)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return true, nil
		}
		return false, c.JSON(200, &utils.Response{
			Success: false,
			Message: "Sign up failed",
			Payload: "",
		})
	}

	return false, c.JSON(200, &utils.Response{
		Success: false,
		Message: "Email already exists",
		Payload: "",
	})
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
		return userInfo, false, c.JSON(200, &utils.Response{
			Success: false,
			Message: "Sign up failed",
			Payload: "",
		})
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
		time.Duration(15)*time.Minute,
	).Err()
	if err != nil {
		return "", false, c.JSON(200, &utils.Response{
			Success: false,
			Message: "Sign up failed",
			Payload: "",
		})
	}
	return verifyEmailToken, true, nil
}

func (controller *SignupController) Execute(c echo.Context, store database.Store) error {
	var req signupRequest
	if err := c.Bind(&req); err != nil {
		return err
	}
	if err := c.Validate(&req); err != nil {
		return err
	}

	ok, err := controller.checkEmailExists(c, store, req)
	if !ok {
		return err
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return c.JSON(200, &utils.Response{
			Success: false,
			Message: "Sign up failed",
			Payload: "",
		})
	}

	userInfo, ok, err := controller.createNewUser(c, store, req, hashedPassword)
	if !ok {
		return err
	}

	_, ok, err = controller.generateVerifyEmailToken(c, userInfo)
	if !ok {
		return err
	}

	return c.JSON(200, &utils.Response{
		Success: true,
		Message: "Sign up success",
		Payload: &signupResponsePayload{
			UserID: userInfo.UserID,
		},
	})
}
