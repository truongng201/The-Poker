package controller

type AppController struct {
	HealthCheckController    HealthCheckController
	SigninController         SigninController
	SignupController         SignupController
	SignoutController        SignoutController
	ResetPasswordController  ResetPasswordController
	ForgotPasswordController ForgotPasswordController
	VerifyEmailController    VerifyEmailController
}
