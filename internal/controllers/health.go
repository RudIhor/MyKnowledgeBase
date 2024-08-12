package controllers

import (
	"net/http"

	"github.com/RivGames/my-knowledge-base/internal/model"
	"github.com/labstack/echo/v4"
)

func Up(c echo.Context) error {
	status := model.Response{
		"status": "OK",
	}

	return c.JSON(http.StatusOK, status)
}
