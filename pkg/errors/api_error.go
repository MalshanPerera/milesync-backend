package errpkg

import (
	"net/http"
)

type ApiError struct {
	Code    int                 `json:"code"`
	Message string              `json:"message,omitempty"`
	Errors  map[string][]string `json:"errors,omitempty"`
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

func UnprocessableEntity(errors map[string][]string) ApiError {
	return ApiError{
		Code:   http.StatusUnprocessableEntity,
		Errors: errors,
	}
}
