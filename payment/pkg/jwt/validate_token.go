package jwt

import (
	"go-todo-app/pkg/configs"

	"github.com/golang-jwt/jwt/v5"
)

func ValidateToken(tokenStr string) (*jwt.Token, error) {
	secret := configs.Load().SECRET

	return jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
		return secret, nil
	})
}