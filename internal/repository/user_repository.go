package repository

import "github.com/RivGames/my-knowledge-base/internal/model"

type UserRepository interface {
	GetUsers() ([]*model.User, error)
	CreateUser() error
}
