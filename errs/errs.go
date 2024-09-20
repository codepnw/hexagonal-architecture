package errs

import "net/http"

type AppError struct {
	Code    int
	Message string
}

// func can return type error, type error need method Error() string
func (e AppError) Error() string {
	return e.Message
}

func NewErrNotFound(message string) error {
	return AppError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

func NewErrUnexpected() error {
	return AppError{
		Code:    http.StatusInternalServerError,
		Message: "unexpected error",
	}
}