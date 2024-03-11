package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// masih mentah baru raba2

func BearerAuth(expectedRole string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")

			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"error": "Unauthorized",
				})
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("JWT_SECRET")), nil
			})

			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"error": "Invalid token",
				})
			}

			if token.Valid {
				// Access the claims
				claims, ok := token.Claims.(jwt.MapClaims)
				if !ok {
					return c.JSON(http.StatusUnauthorized, map[string]interface{}{
						"error": "Invalid token claims",
					})
				}

				if role, ok := claims["role"].(string); ok {
					if role != expectedRole {
						return c.JSON(http.StatusUnauthorized, map[string]interface{}{
							"error": "Unauthorized",
						})
					}
				} else {
					return c.JSON(http.StatusUnauthorized, map[string]interface{}{
						"error": "Invalid role in token claims",
					})
				}

				return next(c)
			}

			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"error": "Invalid token",
			})
		}
	}
}
