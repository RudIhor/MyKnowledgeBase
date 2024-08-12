package repository

import (
	"github.com/RivGames/my-knowledge-base/internal/model"
	"github.com/RivGames/my-knowledge-base/internal/request"
	"gorm.io/gorm"
)

type AnswerRepository interface {
	FetchAll() ([]model.Answer, error)
	Create(*request.CreateAnswerRequest) (*model.Answer, error)
	FetchByID(uint) (*model.Answer, error)
	FetchAnswersByQuestionID(uint) ([]model.Answer, error)
	Update(*model.Answer, *request.UpdateAnswerRequest) (*model.Answer, error)
	Delete(*model.Answer) error
}

type AnswerRepo struct {
	db *gorm.DB
}

func NewAnswerRepository(db *gorm.DB) AnswerRepository {
	return &AnswerRepo{db: db}
}

func (r *AnswerRepo) FetchAll() ([]model.Answer, error) {
	var answers []model.Answer
	if err := r.db.Find(&answers).Error; err != nil {
		return nil, err
	}

	return answers, nil
}

func (r *AnswerRepo) FetchByID(id uint) (*model.Answer, error) {
	var answer *model.Answer
	if err := r.db.First(&answer, id).Error; err != nil {
		return nil, err
	}

	return answer, nil
}

func (r *AnswerRepo) FetchAnswersByQuestionID(questionID uint) ([]model.Answer, error) {
	var answers []model.Answer
	if err := r.db.Where("question_id = ?", questionID).Find(&answers).Error; err != nil {
		return nil, err
	}

	return answers, nil
}

func (r *AnswerRepo) Create(req *request.CreateAnswerRequest) (*model.Answer, error) {
	answer := &model.Answer{
		CreateAnswerRequest: *req,
	}

	return answer, r.db.Create(answer).Error
}

func (r *AnswerRepo) Update(answer *model.Answer, req *request.UpdateAnswerRequest) (*model.Answer, error) {
	answer.Text = req.Text

	if err := r.db.Save(&answer).Error; err != nil {
		return nil, err
	}

	return answer, nil
}

func (r *AnswerRepo) Delete(answer *model.Answer) error {
	if err := r.db.Delete(answer).Error; err != nil {
		return err
	}

	return nil
}
