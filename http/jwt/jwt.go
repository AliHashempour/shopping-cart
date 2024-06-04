package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var SecretKey = []byte("secret-key")

type Claims struct {
	jwt.RegisteredClaims
	UserId uint
}

func GenerateToken(userId uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userId": userId,
			"exp":    time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
