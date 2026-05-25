package jwt

import (
	"fmt"
	"gateway/pkg/configs"

	"github.com/golang-jwt/jwt/v5"
)

func ValidateToken(tokenStr string) (*jwt.Token, error) {
	secret := configs.Load().SECRET

return jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method")
	}
	return []byte(secret), nil
})
}