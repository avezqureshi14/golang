package routes

import (
	"go-todo-app/internal/auth/handler"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine, h *handler.AuthHandler, jwt gin.HandlerFunc) {
	auth := r.Group("/auth")

	auth.POST("/signup", h.Signup)
	auth.POST("/signin", h.Signin)

	protected := r.Group("/users")
	protected.Use(jwt)
	protected.GET("", h.GetAllUsers)
}