package controller

import (
	database "auth-service/pkg/database"
)

type AppController struct {
	HealthCheckController    HealthCheckController
	SigninController         SigninController
	SignupController         SignupController
	SignoutController        SignoutController
	ResetPasswordController  ResetPasswordController
	ForgotPasswordController ForgotPasswordController
	VerifyEmailController    VerifyEmailController
	Store                    database.Store
}
