package middleware

import (
	jwtUtil "basket/http/jwt" // Assuming this is the correct path
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authorizationHeader := c.Request().Header.Get("Authorization")
		if authorizationHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing Authorization header")
		}
		tokenString := strings.TrimPrefix(authorizationHeader, "Bearer ")
		if tokenString == authorizationHeader {
			return echo.NewHTTPError(http.StatusUnauthorized, "Bearer token not found")
		}

		claims := &jwtUtil.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, echo.NewHTTPError(http.StatusUnauthorized, "Unexpected signing method")
			}
			return []byte("jwt-secret"), nil
		})

		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or expired JWT")
		}
		if token.Valid {
			c.Set("userId", claims.UserId)
			return next(c)
		} else {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid JWT claims")
		}
	}
}
