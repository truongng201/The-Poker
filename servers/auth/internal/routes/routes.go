package routes

import (
	controller "auth-service/internal/controller"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo, controller *controller.AppController) *echo.Echo {
	// ApiRoutes
	api := e.Group("/auth")

	api.GET("/health", func(c echo.Context) error {
		return controller.HealthCheckController.Execute(c, controller.Store)
	})

	api.POST("/signin", func(c echo.Context) error {
		return controller.SigninController.Execute(c, controller.Store)
	})

	api.POST("/signup", func(c echo.Context) error {
		return controller.SignupController.Execute(c, controller.Store)
	})

	api.POST("/signout", func(c echo.Context) error {
		return controller.SignoutController.Execute(c)
	})

	api.POST("/reset-password", func(c echo.Context) error {
		return controller.ResetPasswordController.Execute(c)
	})

	api.POST("/forgot-password", func(c echo.Context) error {
		return controller.ForgotPasswordController.Execute(c, controller.Store)
	})

	api.POST("/verify-email", func(c echo.Context) error {
		return controller.VerifyEmailController.Execute(c)
	})

	return e

}
