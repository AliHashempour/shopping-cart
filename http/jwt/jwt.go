package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtSecret = "jwt-secret"

type Claims struct {
	jwt.RegisteredClaims
	UserId uint `json:"user_id"`
}

func GenerateToken(userId uint) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := Claims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "basket-app",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}
