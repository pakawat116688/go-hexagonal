package errs

import "net/http"

type AppError struct {
	Code int
	Message string
}

func (app AppError) Error() string  {
	return app.Message
}

func NewNotfoundError(message string) error {
	return AppError{
		Code: http.StatusNotFound,
		Message: message,
	}
}

func NewUnexpectedError() error {
	return AppError{
		Code: http.StatusNotFound,
		Message: "unexpected error",
	}
}

func NewDeleteError() error {
	return AppError{
		Code: http.StatusNotFound,
		Message: "cannot delete row",
	}
}

func NewValidationError() error  {
	return AppError{
		Code: http.StatusUnprocessableEntity,
		Message: "request body incorrect",
	}
}