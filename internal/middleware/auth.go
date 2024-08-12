package middleware

import (
	"net/http"
	"strings"

	"github.com/RivGames/my-knowledge-base/pkg/errs"
	"github.com/RivGames/my-knowledge-base/pkg/jwt"
	"github.com/labstack/echo/v4"
)

func WithAuthentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := strings.TrimPrefix(c.Request().Header.Get("Authorization"), "Bearer ")
		_, err := jwt.VerifyToken(token)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, errs.ErrInvalidToken.Error())
		}
		if err := next(c); err != nil {
			c.Error(err)
		}
		return nil
	}
}
