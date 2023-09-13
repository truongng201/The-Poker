package routes

import (
	"log"
	"net/http"

	controller "auth-service/internal/controller"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	log.Println(i)
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func Routes(e *echo.Echo, controller controller.AppController) *echo.Echo {
	e.Validator = &CustomValidator{validator: validator.New()}
	// ApiRoutes
	api := e.Group("/auth")

	api.GET("/health", func(c echo.Context) error {
		return controller.HealthCheckController.Execute(c)
	})

	api.POST("/signin", func(c echo.Context) error {
		return controller.SigninController.Execute(c)
	})

	api.POST("/signup", func(c echo.Context) error {
		return controller.SignupController.Execute(c)
	})

	api.POST("/signout", func(c echo.Context) error {
		return controller.SignoutController.Execute(c)
	})

	api.POST("/reset-password", func(c echo.Context) error {
		return controller.ResetPasswordController.Execute(c)
	})

	api.POST("/forgot-password", func(c echo.Context) error {
		return controller.ForgotPasswordController.Execute(c)
	})

	api.POST("/verify-email", func(c echo.Context) error {
		return controller.VerifyEmailController.Execute(c)
	})

	return e

}
