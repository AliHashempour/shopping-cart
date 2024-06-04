package middleware

import (
	jwtUtil "basket/http/jwt"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Authorization header is required")
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			return echo.NewHTTPError(http.StatusUnauthorized, "Bearer token not found")
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtUtil.SecretKey, nil
		})

		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": err.Error()})
		}

		if !token.Valid {
			return fmt.Errorf("invalid token")
		}

		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid JWT claims")
	}
}
