package errors

import (
	"net/http"
)

type ApiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e ApiError) Error() string {
	return e.Message
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

func (e ApiError) UnprocessableEntity(message string) ApiError {
	return ApiError{
		Code:    http.StatusUnprocessableEntity,
		Message: message,
	}
}
