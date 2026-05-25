package routes

import (
	handler "go-todo-app/internal/auth/handler"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine, h *handler.AuthHandler, jwt gin.HandlerFunc) {
	r.POST("/signup", h.Signup)
	r.POST("/signin", h.Signin)
	r.GET("/", h.GetAllUsers)
}
