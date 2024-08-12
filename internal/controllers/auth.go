package controllers

import (
	"net/http"

	"github.com/RivGames/my-knowledge-base/internal/model"
	"github.com/RivGames/my-knowledge-base/internal/repository"
	"github.com/RivGames/my-knowledge-base/internal/request"
	"github.com/RivGames/my-knowledge-base/internal/service"
	"github.com/RivGames/my-knowledge-base/pkg/errs"
	"github.com/RivGames/my-knowledge-base/pkg/jwt"
	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	cc := c.(*model.CustomContext)

	registerRequest := new(request.RegisterUserRequest)
	if err := c.Bind(registerRequest); err != nil {
		return echo.NewHTTPError(errs.ErrUnableToBindRequest.StatusCode, errs.ErrUnableToBindRequest.Error())
	}
	authService := service.NewAuthService(repository.NewUserRepository(cc.App.Store.DB))
	user, err := authService.Register(c, registerRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	return c.JSON(http.StatusCreated, user)
}

func Login(c echo.Context) error {
	cc := c.(*model.CustomContext)

	request := new(request.LoginUserRequest)
	if err := c.Bind(request); err != nil {
		return echo.NewHTTPError(errs.ErrUnableToBindRequest.StatusCode, errs.ErrUnableToBindRequest.Error())
	}
	authService := service.NewAuthService(repository.NewUserRepository(cc.App.Store.DB))
	user, err := authService.Login(c, request)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	return c.JSON(http.StatusOK, user)
}

func Me(c echo.Context) error {
	cc := c.(*model.CustomContext)

	userId, err := jwt.GetUserID(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	user, err := repository.NewUserRepository(cc.App.Store.DB).FetchById(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}
