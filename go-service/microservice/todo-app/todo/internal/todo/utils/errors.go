package todo

import (
	"go-todo-app/pkg/errors"
	"net/http"
)

var (
	ErrTitleTooLong = errors.AppError{
		Code:    "TITLE_TOO_LONG",
		Message: "title exceeds allowed length",
		Status:  http.StatusBadRequest,
	}
)
