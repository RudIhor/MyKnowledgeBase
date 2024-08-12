package service

import (
	"net/http"

	"github.com/RivGames/my-knowledge-base/internal/model"
	"github.com/RivGames/my-knowledge-base/internal/repository"
	"github.com/RivGames/my-knowledge-base/internal/request"
	"github.com/RivGames/my-knowledge-base/pkg/errs"
	"github.com/RivGames/my-knowledge-base/pkg/jwt"
	"github.com/labstack/echo/v4"
)

type QuestionService struct {
	questionRepo repository.QuestionRepository
}

func NewQuestionService(questionRepo repository.QuestionRepository) *QuestionService {
	return &QuestionService{questionRepo: questionRepo}
}

func (s *QuestionService) GetAllQuestions() ([]model.Question, error) {
	return s.questionRepo.FetchAll()
}

func (s *QuestionService) CreateQuestion(c echo.Context, req *request.CreateQuesitonRequest) (*model.Question, error) {
	if err := c.Validate(req); err != nil {
		return nil, err
	}
	return s.questionRepo.Create(req)
}

func (s *QuestionService) GetQuestionByID(id int) (*model.Question, error) {
	return s.questionRepo.FetchByID(id)
}

func (s *QuestionService) GetQuestionByIDAndUserID(c echo.Context, id int) (*model.Question, error) {
	question, err := s.questionRepo.FetchByID(id)
	if err != nil {
		return nil, err
	}
	userID, _ := jwt.GetUserID(c)
	if question.UserId != userID {
		return nil, errs.ErrEntityDoesNotBelongsToUser
	}

	return question, nil
}

func (s *QuestionService) UpdateQuestionByID(c echo.Context, id int, req *request.UpdateQuestionRequest) (*model.Question, error) {
	question, err := s.GetQuestionByIDAndUserID(c, id)
	if err != nil {
		return nil, errs.ErrEntityDoesNotBelongsToUser
	}
	if err := c.Validate(req); err != nil {
		return nil, errs.NewAPIError(http.StatusUnprocessableEntity, err.Error())
	}

	return s.questionRepo.Update(question, req)
}

func (s *QuestionService) DeleteQuestionByID(c echo.Context, id int) error {
	question, err := s.GetQuestionByIDAndUserID(c, id)
	if err != nil {
		return errs.ErrEntityDoesNotBelongsToUser
	}

	return s.questionRepo.Delete(question)
}
