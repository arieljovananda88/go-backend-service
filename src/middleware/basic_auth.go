package middleware

import (
	"crypto/subtle"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func BasicAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			providedUsername, providedPassword, hasAuth := ctx.Request().BasicAuth()

			if hasAuth &&
				subtle.ConstantTimeCompare([]byte(providedUsername), []byte(os.Getenv("BASIC_AUTH_USERNAME"))) == 1 &&
				subtle.ConstantTimeCompare([]byte(providedPassword), []byte(os.Getenv("BASIC_AUTH_PASSWORD"))) == 1 {
				return next(ctx)
			}

			print(providedPassword)
			ctx.Response().Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			return echo.NewHTTPError(http.StatusUnauthorized, providedPassword)
		}
	}
}
