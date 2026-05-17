package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte("super-secret-key")

type Claims struct {
	Service string `json:"service"`
	jwt.RegisteredClaims
}

func GenerateServiceToken(serviceName string) (string, error) {
	claims := Claims{
		Service: serviceName,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func ValidateToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (any, error) {
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)

	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}
