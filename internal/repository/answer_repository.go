package repository

import "gorm.io/gorm"

type AnswerRepository interface {
}

type AnswerRepo struct {
	db *gorm.DB
}

func NewAnswerRepository(db *gorm.DB) AnswerRepository {
	return &AnswerRepo{db: db}
}
