package model

import (
	"github.com/RivGames/my-knowledge-base/pkg/app"
	"github.com/labstack/echo/v4"
)

type CustomContext struct {
	echo.Context
	App *app.App
}
