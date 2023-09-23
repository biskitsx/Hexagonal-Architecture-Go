package errs

import "net/http"

type AppError struct {
	Code    int
	Message string
}

func (e AppError) Error() string {
	return e.Message
}

func NewNotfoundError(msg string) error {
	return AppError{
		Code:    http.StatusNotFound,
		Message: msg,
	}
}
func NewUnexpectError() error {
	return AppError{
		Code:    http.StatusInternalServerError,
		Message: "unexpect error",
	}
}

func NewValidationError(message string) error {
	return AppError{
		Code:    http.StatusUnprocessableEntity,
		Message: message,
	}
}
