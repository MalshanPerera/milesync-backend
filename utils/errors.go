package utils

import "net/http"

type AppError struct {
	Message string
}

type ApiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e ApiError) Error() string {
	return e.Message
}

func (e AppError) Error() string {
	return e.Message
}

var BadRequest = ApiError{
	Code:    http.StatusBadRequest,
	Message: "Bad request",
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
