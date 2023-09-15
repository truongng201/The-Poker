package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload"`
}

const (
	ErrBadRequest         = "Bad request"
	ErrEmailAlreadyExists = "User with given email already exists"
	ErrNoSuchUser         = "User not found"
	ErrWrongCredentials   = "Wrong Credentials"
	ErrNotFound           = "Not Found"
	ErrUnauthorized       = "Unauthorized"
	ErrForbidden          = "Forbidden"
	ErrBadQueryParams     = "Invalid query params"
	ErrInternalServer     = "Internal server error"
	ErrUnverifiedEmail    = "Email not verified"
)

func ErrInternalServerRepsonse() error {
	return echo.NewHTTPError(http.StatusInternalServerError, Response{
		Success: false,
		Message: ErrInternalServer,
		Payload: "",
	})
}

func ErrBadRequestResponse() error {
	return echo.NewHTTPError(http.StatusBadRequest, Response{
		Success: false,
		Message: ErrBadRequest,
		Payload: "",
	})
}

func ErrEmailAlreadyExistsResponse() error {
	return echo.NewHTTPError(http.StatusBadRequest, Response{
		Success: false,
		Message: ErrEmailAlreadyExists,
		Payload: "",
	})
}

func ErrNoSuchUserResponse() error {
	return echo.NewHTTPError(http.StatusNotFound, Response{
		Success: false,
		Message: ErrNoSuchUser,
		Payload: "",
	})
}

func ErrWrongCredentialsResponse() error {
	return echo.NewHTTPError(http.StatusUnauthorized, Response{
		Success: false,
		Message: ErrWrongCredentials,
		Payload: "",
	})
}

func ErrNotFoundResponse() error {
	return echo.NewHTTPError(http.StatusNotFound, Response{
		Success: false,
		Message: ErrNotFound,
		Payload: "",
	})
}

func ErrUnauthorizedResponse() error {
	return echo.NewHTTPError(http.StatusUnauthorized, Response{
		Success: false,
		Message: ErrUnauthorized,
		Payload: "",
	})
}

func ErrForbiddenResponse() error {
	return echo.NewHTTPError(http.StatusForbidden, Response{
		Success: false,
		Message: ErrForbidden,
		Payload: "",
	})
}

func ErrBadQueryParamsResponse() error {
	return echo.NewHTTPError(http.StatusBadRequest, Response{
		Success: false,
		Message: ErrBadQueryParams,
		Payload: "",
	})
}

func ErrInternalServerResponse() error {
	return echo.NewHTTPError(http.StatusInternalServerError, Response{
		Success: false,
		Message: ErrInternalServer,
		Payload: "",
	})
}

func ErrUnverifiedEmailResponse() error {
	return echo.NewHTTPError(http.StatusBadRequest, Response{
		Success: false,
		Message: ErrUnverifiedEmail,
		Payload: "",
	})
}