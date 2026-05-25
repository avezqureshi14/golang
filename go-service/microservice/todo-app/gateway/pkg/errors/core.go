package errors

import "net/http"

type AppError struct {
	Code    string
	Message string
	Status  int
}

func (e AppError) Error() string {
	return e.Message
}


var (
	ErrNotFound = AppError{
		Code:    "NOT_FOUND",
		Message: "resource not found",
		Status:  http.StatusNotFound,
	}

	ErrInvalidInput = AppError{
		Code:    "INVALID_INPUT",
		Message: "invalid input",
		Status:  http.StatusBadRequest,
	}
)