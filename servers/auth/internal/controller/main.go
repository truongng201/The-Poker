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

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload"`
}