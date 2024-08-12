package controllers

import (
	"net/http"

	"github.com/RivGames/my-knowledge-base/internal/model"
	"github.com/RivGames/my-knowledge-base/internal/repository"
	"github.com/RivGames/my-knowledge-base/internal/request"
	"github.com/RivGames/my-knowledge-base/internal/service"
	"github.com/RivGames/my-knowledge-base/pkg/errs"
	"github.com/RivGames/my-knowledge-base/pkg/helpers"
	"github.com/RivGames/my-knowledge-base/pkg/jwt"
	"github.com/labstack/echo/v4"
)

func GetQuestions(c echo.Context) error {
	cc := c.(*model.CustomContext)

	questionRepository := repository.NewQuesitonRepository(cc.App.Store.DB)
	questionService := service.NewQuestionService(questionRepository)
	questions, err := questionService.GetAllQuestions()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, questions)
}

func CreateQuestion(c echo.Context) error {
	cc := c.(*model.CustomContext)

	createQuestionRequest := new(request.CreateQuesitonRequest)
	if err := c.Bind(createQuestionRequest); err != nil {
		return echo.NewHTTPError(errs.ErrUnableToBindRequest.StatusCode, errs.ErrUnableToBindRequest.Error())
	}
	userID, err := jwt.GetUserID(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	createQuestionRequest.UserId = userID

	questionRepository := repository.NewQuesitonRepository(cc.App.Store.DB)
	questionService := service.NewQuestionService(questionRepository)

	question, err := questionService.CreateQuestion(c, createQuestionRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, question)
}

func GetQuestion(c echo.Context) error {
	cc := c.(*model.CustomContext)

	id, err := helpers.GetIDFromParam(c.Param("id"))
	questionRepository := repository.NewQuesitonRepository(cc.App.Store.DB)
	questionService := service.NewQuestionService(questionRepository)
	question, err := questionService.GetQuestionByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, question)
}

func UpdateQuestion(c echo.Context) error {
	cc := c.(*model.CustomContext)

	updateRequest := new(request.UpdateQuestionRequest)
	if err := c.Bind(updateRequest); err != nil {
		return echo.NewHTTPError(errs.ErrUnableToBindRequest.StatusCode, errs.ErrUnableToBindRequest.Error())
	}
	questionRepository := repository.NewQuesitonRepository(cc.App.Store.DB)
	questionService := service.NewQuestionService(questionRepository)
	id, _ := helpers.GetIDFromParam(c.Param("id"))
	question, err := questionService.UpdateQuestionByID(c, id, updateRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, question)
}

func DeleteQuestion(c echo.Context) error {
	cc := c.(*model.CustomContext)

	questionRepository := repository.NewQuesitonRepository(cc.App.Store.DB)
	questionService := service.NewQuestionService(questionRepository)
	id, _ := helpers.GetIDFromParam(c.Param("id"))
	if err := questionService.DeleteQuestionByID(c, id); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusNoContent, "")
}
