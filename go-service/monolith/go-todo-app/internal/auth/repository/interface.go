package auth

import (
	models "go-todo-app/internal/auth/models"
)

type AuthRepository interface {
	Create(user models.User) (models.User, error)
	GetByEmail(email string) (models.User, error)
	GetAll() ([]models.User, error)
}