package errpkg

import "github.com/jackc/pgx/v5"

type AppError struct {
	Message string
}

type DBError struct {
	Message string
}

func NewDBError(message string) DBError {
	return DBError{
		Message: message,
	}
}

func (e DBError) Error() string {
	return e.Message
}

func NewAppError(message string) AppError {
	return AppError{
		Message: message,
	}
}

func ServiceError(message string) error {
	return NewAppError(message)
}

func (e AppError) Error() string {
	return e.Message
}

var NoResults = pgx.ErrNoRows
