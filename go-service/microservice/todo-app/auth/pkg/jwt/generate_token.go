package jwt

import (
	"go-todo-app/pkg/configs"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID int, email string) (string, error) {
	secret := configs.Load().SECRET
	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
