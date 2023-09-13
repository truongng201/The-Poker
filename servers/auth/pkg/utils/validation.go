package utils

import (
	"auth-service/internal/controller"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		log.Error(err.Error())
		return echo.NewHTTPError(http.StatusOK, &controller.Response{
			Success: false,
			Message: "Validation error",
			Payload: err.Error(),
		})
	}
	return nil
}