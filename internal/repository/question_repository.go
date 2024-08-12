package repository

import (
	"github.com/RivGames/my-knowledge-base/internal/model"
	"github.com/RivGames/my-knowledge-base/internal/request"
	"gorm.io/gorm"
)

type QuestionRepository interface {
	FetchAll() ([]model.Question, error)
	Create(*request.CreateQuesitonRequest) (*model.Question, error)
	FetchByID(int) (*model.Question, error)
	Update(*model.Question, *request.UpdateQuestionRequest) (*model.Question, error)
	Delete(*model.Question) error
}

type QuestionRepo struct {
	db *gorm.DB
}

func NewQuesitonRepository(db *gorm.DB) QuestionRepository {
	return &QuestionRepo{db: db}
}

func (r *QuestionRepo) FetchAll() ([]model.Question, error) {
	var questions []model.Question
	if err := r.db.Find(&questions).Error; err != nil {
		return nil, err
	}

	return questions, nil
}

func (r *QuestionRepo) FetchByID(id int) (*model.Question, error) {
	var question *model.Question
	if err := r.db.First(&question, id).Error; err != nil {
		return nil, err
	}

	return question, nil
}

func (r *QuestionRepo) Create(req *request.CreateQuesitonRequest) (*model.Question, error) {
	question := &model.Question{
		CreateQuesitonRequest: *req,
	}

	return question, r.db.Create(question).Error
}

func (r *QuestionRepo) Update(question *model.Question, req *request.UpdateQuestionRequest) (*model.Question, error) {
	question.Title = req.Title
	if err := r.db.Save(&question).Error; err != nil {
		return nil, err
	}

	return question, nil
}

func (r *QuestionRepo) Delete(question *model.Question) error {
	if err := r.db.Delete(question).Error; err != nil {
		return err
	}

	return nil
}
