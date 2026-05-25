package http_response

import (
	"go-todo-app/pkg/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)


type SuccessResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

func SendSuccess(c *gin.Context, status int, data interface{}) {
	c.JSON(status, SuccessResponse{
		Success: true,
		Data:    data,
	})
}

func SendError(c *gin.Context, status int, message string) {
	c.JSON(status, ErrorResponse{
		Success: false,
		Error:   message,
	})
}

func HandleError(c *gin.Context, err error) {
	if appErr, ok := err.(errors.AppError); ok {
		SendError(c, appErr.Status, appErr.Message)
		return
	}

	SendError(c, http.StatusInternalServerError, "internal server error")
}