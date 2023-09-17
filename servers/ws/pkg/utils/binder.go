package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type CustomBinder struct{}

func (cb *CustomBinder) Bind(i interface{}, c echo.Context) error {
	db := new(echo.DefaultBinder)
	if err := db.Bind(i, c); err != nil {
		log.Error(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, Response{
			Success: false,
			Message: "Invalid request",
			Payload: err.Error(),
		})

	}
	return nil
}
