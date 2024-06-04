package middleware

import (
	jwtUtil "basket/http/jwt"
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

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if userId, ok := claims["sub"].(string); ok {
				c.Set("userId", userId)
				return next(c)
			}
		}

		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid JWT claims")

	}
}
