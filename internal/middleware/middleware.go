package middleware

import (
	"github.com/RivGames/my-knowledge-base/pkg/app"
	"github.com/labstack/echo/v4"
)

type Middleware struct {
	*echo.Echo
	App *app.App
}

func New(echo *echo.Echo, app *app.App) *Middleware {
	m := &Middleware{Echo: echo, App: app}

	m.Use(m.RegisterCustomContext())

	return m
}
