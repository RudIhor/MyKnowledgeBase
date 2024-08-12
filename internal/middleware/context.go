package middleware

import (
	"github.com/RivGames/my-knowledge-base/internal/model"
	"github.com/labstack/echo/v4"
)

func (m *Middleware) RegisterCustomContext() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &model.CustomContext{
				Context: c,
				App:     m.App,
			}

			return next(cc)
		}
	}
}
