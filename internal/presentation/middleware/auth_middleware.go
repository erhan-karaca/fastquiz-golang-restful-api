package middleware

import (
	"fastquiz-api/pkg/config"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		authHeader = "Bearer 12345" // TODO:: Chrome Test için manuel eklendi
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "missing authorization header",
			})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "invalid authorization header format",
			})
		}

		token := parts[1]

		if !validateToken(token) {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "invalid token",
			})
		}

		return next(c)
	}
}

func validateToken(token string) bool {
	return token == config.AppConfig.FrontEndToken
}
