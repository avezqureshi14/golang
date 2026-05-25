package handler

import (
	"log"
	"net/http"

	dto "go-todo-app/internal/auth/dto"
	service "go-todo-app/internal/auth/service"
	"go-todo-app/pkg/jwt"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(s *service.AuthService) *AuthHandler {
	return &AuthHandler{service: s}
}

func (h *AuthHandler) Signup(c *gin.Context) {
	var req dto.SignupRequest
	c.ShouldBindJSON(&req)

	user, err := h.service.Signup(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := jwt.GenerateToken(user.ID, user.Email)
	if err != nil {
		log.Print("JWT error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "token generation failed"})
		return
	}
	c.JSON(http.StatusOK, dto.AuthResponse{Token: token})
}

func (h *AuthHandler) Signin(c *gin.Context) {
	var req dto.SigninRequest
	c.ShouldBindJSON(&req)

	user, err := h.service.Signin(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := jwt.GenerateToken(user.ID, user.Email)
	if err != nil {
		log.Print("JWT error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "token generation failed"})
		return
	}

	c.JSON(http.StatusOK, dto.AuthResponse{Token: token})
}

func (h *AuthHandler) GetAllUsers(c *gin.Context) {
	users, _ := h.service.GetAllUsers()
	c.JSON(http.StatusOK, users)
}
