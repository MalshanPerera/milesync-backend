package common

import (
	"net/http"

	"github.com/jackc/pgx/v5"
)

type AppError struct {
	Message string
}

type ApiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type DBError struct {
	Message string
}

var NoResults = pgx.ErrNoRows

func (e ApiError) Error() string {
	return e.Message
}

func (e AppError) Error() string {
	return e.Message
}

func (e DBError) Error() string {
	return e.Message
}

func NewAppError(message string) AppError {
	return AppError{
		Message: message,
	}
}

func NewDBError(message string) DBError {
	return DBError{
		Message: message,
	}
}

func NewApiError(code int, message string) ApiError {
	return ApiError{
		Code:    code,
		Message: message,
	}
}

func BadRequest(message string) ApiError {
	return ApiError{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}

var InternalServerError = ApiError{
	Code:    http.StatusInternalServerError,
	Message: "Internal server error",
}

func (e ApiError) UnprocessableEntity(message string) ApiError {
	return ApiError{
		Code:    http.StatusUnprocessableEntity,
		Message: message,
	}
}
