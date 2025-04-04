package utils

import (
	"os"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	ID uint
	jwt.RegisteredClaims
}

func GenerateJWT(payload JWTClaims) (string, error) {
	claims := JWTClaims{
		ID: payload.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	return token.SignedString([]byte(secret))
}