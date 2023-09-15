package controller

import (
	"fmt"
	"time"

	config "auth-service/config"
	database "auth-service/pkg/database"
	sqlc "auth-service/pkg/database/sqlc"
	utils "auth-service/pkg/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type signinRequestBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32,alphanum"`
}

type signinResponsePayload struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

type SigninController struct{}

func (controller *SigninController) checkEmailExists(
	c echo.Context,
	store database.Store,
	req signinRequestBody,
) (sqlc.GetUserByEmailRow, bool, error) {
	res, err := store.GetUserByEmail(c.Request().Context(), req.Email)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return res, false, utils.ErrNotFoundResponse()
		}
		log.Error(err)
		return res, false, utils.ErrInternalServerRepsonse()
	}

	if !res.IsVerified {
		return res, false, c.Redirect(302, fmt.Sprintf("%s/reverify", config.Con.Domains.Client))
	}

	return res, true, nil
}

func (controller *SigninController) checkPassword(
	c echo.Context,
	req signinRequestBody,
	res sqlc.GetUserByEmailRow,
) (bool, error) {
	if !utils.CheckPassword(req.Password, res.HashedPassword) {
		return false, utils.ErrWrongCredentialsResponse()
	}

	return true, nil
}

func (controller *SigninController) generateNewRefreshToken(
	c echo.Context,
	store database.Store,
	userInfo sqlc.GetUserByEmailRow,
) (string, bool, error) {
	newRefreshToken, err := utils.GenerateJWT(&jwt.MapClaims{
		"sub": map[string]interface{}{
			"user_id":  userInfo.UserID,
			"username": userInfo.Username,
			"email":    userInfo.Email,
		},
		"iat": time.Now().Local(),
		"exp": time.Duration(config.Con.JWT.RefreshTokenExpirationTime) * time.Minute,
	}, config.Con.JWT.SecretKey)

	if err != nil {
		log.Error(err)
		return "", false, utils.ErrInternalServerRepsonse()
	}

	_, err = store.CreateRefreshToken(
		c.Request().Context(),
		sqlc.CreateRefreshTokenParams{
			UserID: userInfo.ID,
			DeviceType: pgtype.Text{
				String: c.Request().Host,
			},
			UserAgent: pgtype.Text{
				String: c.Request().UserAgent(),
			},
			IpAddress: c.Request().RemoteAddr,
			Token:     newRefreshToken,
		},
	)

	if err != nil {
		log.Error(err)
		return "", false, utils.ErrInternalServerRepsonse()
	}

	return newRefreshToken, true, nil
}

func (controller *SigninController) generateNewAccessToken(
	c echo.Context,
	store database.Store,
	userInfo sqlc.GetUserByEmailRow,
) (string, bool, error) {
	newAccessToken, err := utils.GenerateJWT(&jwt.MapClaims{
		"iat": time.Now().Local(),
		"exp": time.Duration(config.Con.JWT.AccessTokenExpirationTime) * time.Minute,
	}, config.Con.JWT.SecretKey)
	if err != nil {
		return "", false, utils.ErrInternalServerRepsonse()
	}

	err = utils.RedisClient.Set(
		c.Request().Context(),
		newAccessToken,
		userInfo.UserID,
		time.Duration(config.Con.JWT.AccessTokenExpirationTime)*time.Minute,
	).Err()

	if err != nil {
		log.Error(err)
		return "", false, utils.ErrInternalServerRepsonse()
	}
	return newAccessToken, true, nil
}

func (controller *SigninController) Execute(c echo.Context, store database.Store) error {
	var req signinRequestBody

	if err := c.Bind(&req); err != nil {
		log.Error(err)
		return utils.ErrBadRequestResponse()
	}

	if err := c.Validate(&req); err != nil {
		log.Error(err)
		return utils.ErrBadRequestResponse()
	}

	userInfo, ok, err := controller.checkEmailExists(c, store, req)
	if !ok {
		return err
	}

	ok, err = controller.checkPassword(c, req, userInfo)
	if !ok {
		return err
	}

	refreshToken, ok, err := controller.generateNewRefreshToken(c, store, userInfo)
	if !ok {
		return err
	}

	accessToken, ok, err := controller.generateNewAccessToken(c, store, userInfo)
	if !ok {
		return err
	}

	return utils.SuccessResponse(
		"Sign in successfully",
		&signinResponsePayload{
			RefreshToken: refreshToken,
			AccessToken:  accessToken,
		},
	)
}
