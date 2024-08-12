package controllers

import (
	"net/http"

	"github.com/RivGames/my-knowledge-base/internal/model"
	"github.com/RivGames/my-knowledge-base/internal/repository"
	"github.com/RivGames/my-knowledge-base/internal/request"
	"github.com/RivGames/my-knowledge-base/internal/service"
	"github.com/RivGames/my-knowledge-base/pkg/helpers"
	"github.com/RivGames/my-knowledge-base/pkg/jwt"
	"github.com/labstack/echo/v4"
)

func GetAnswers(c echo.Context) error {
	cc := c.(*model.CustomContext)

	answerService := service.NewAnswerService(repository.NewAnswerRepository(cc.App.Store.DB))

	answers, err := answerService.GetAnswers(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	response := model.Response{
		"data": answers,
	}

	return c.JSON(http.StatusOK, response)
}

func CreateAnswer(c echo.Context) error {
	cc := c.(*model.CustomContext)

	createAnswerRequest := new(request.CreateAnswerRequest)
	if err := c.Bind(createAnswerRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	userID, err := jwt.GetUserID(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	createAnswerRequest.UserId = userID

	answerService := service.NewAnswerService(repository.NewAnswerRepository(cc.App.Store.DB))

	if _, err := repository.NewQuesitonRepository(cc.App.Store.DB).FetchByID(createAnswerRequest.QuestionId); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	answer, err := answerService.CreateAnswer(c, createAnswerRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, answer)
}

func GetAnswer(c echo.Context) error {
	cc := c.(*model.CustomContext)

	answerService := service.NewAnswerService(repository.NewAnswerRepository(cc.App.Store.DB))

	id, _ := helpers.GetIDFromParam(c.Param("id"))

	answer, err := answerService.GetAnswerByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, answer)
}

func GetQuestionAnswers(c echo.Context) error {
	cc := c.(*model.CustomContext)

	answerService := service.NewAnswerService(repository.NewAnswerRepository(cc.App.Store.DB))

	questionID, _ := helpers.GetIDFromParam(c.Param("questionID"))
	question, err := repository.NewQuesitonRepository(cc.App.Store.DB).FetchByID(questionID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	answers, err := answerService.GetQuestionAnswers(question.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	response := model.Response{
		"data": answers,
	}

	return c.JSON(http.StatusOK, response)
}

func UpdateAnswer(c echo.Context) error {
	cc := c.(*model.CustomContext)

	answerService := service.NewAnswerService(repository.NewAnswerRepository(cc.App.Store.DB))

	updateRequest := new(request.UpdateAnswerRequest)
	if err := c.Bind(updateRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	id, _ := helpers.GetIDFromParam(c.Param("id"))
	answer, err := answerService.UpdateAnswerByID(c, id, updateRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, answer)
}

func DeleteAnswer(c echo.Context) error {
	cc := c.(*model.CustomContext)

	answerService := service.NewAnswerService(repository.NewAnswerRepository(cc.App.Store.DB))

	id, _ := helpers.GetIDFromParam(c.Param("id"))

	if err := answerService.DeleteAnswerByID(c, id); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusNoContent, "")
}
