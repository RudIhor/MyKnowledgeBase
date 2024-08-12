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

type AnswerService struct {
	answerRepo repository.AnswerRepository
}

func NewAnswerService(answerRepo repository.AnswerRepository) *AnswerService {
	return &AnswerService{answerRepo: answerRepo}
}

func (s *AnswerService) GetAnswers(c echo.Context) ([]model.Answer, error) {
	return s.answerRepo.FetchAll()
}

func (s *AnswerService) GetQuestionAnswers(questionID uint) ([]model.Answer, error) {
	return s.answerRepo.FetchAnswersByQuestionID(questionID)
}

func (s *AnswerService) CreateAnswer(c echo.Context, req *request.CreateAnswerRequest) (*model.Answer, error) {
	if err := c.Validate(req); err != nil {
		return nil, err
	}

	return s.answerRepo.Create(req)
}

func (s *AnswerService) GetAnswerByID(id uint) (*model.Answer, error) {
	return s.answerRepo.FetchByID(id)
}

func (s *AnswerService) GetAnswerByIDAndUserID(c echo.Context, id uint) (*model.Answer, error) {
	answer, err := s.GetAnswerByID(id)
	if err != nil {
		return nil, err
	}
	userID, _ := jwt.GetUserID(c)
	if answer.UserId != userID {
		return nil, errs.ErrEntityDoesNotBelongsToUser
	}

	return answer, nil
}

func (s *AnswerService) UpdateAnswerByID(c echo.Context, id uint, req *request.UpdateAnswerRequest) (*model.Answer, error) {
	answer, err := s.GetAnswerByIDAndUserID(c, id)
	if err != nil {
		return nil, err
	}
	if err := c.Validate(req); err != nil {
		return nil, errs.NewAPIError(http.StatusUnprocessableEntity, err.Error())
	}

	return s.answerRepo.Update(answer, req)
}

func (s *AnswerService) DeleteAnswerByID(c echo.Context, id uint) error {
	answer, err := s.GetAnswerByIDAndUserID(c, id)
	if err != nil {
		return errs.ErrEntityDoesNotBelongsToUser
	}

	return s.answerRepo.Delete(answer)
}
