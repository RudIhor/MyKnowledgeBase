package repository

import (
	"github.com/RivGames/my-knowledge-base/internal/model"
	requests "github.com/RivGames/my-knowledge-base/internal/requests/auth"
	"gorm.io/gorm"
)

type UserRepository interface {
	Fetch() ([]*model.User, error)
	GetByEmail(string) (*model.User, error)
	Create(*requests.RegisterUserRequest) (*model.User, error)
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepo{db: db}
}

func (u *UserRepo) Fetch() ([]*model.User, error) {
	return nil, nil
}

func (u *UserRepo) Create(request *requests.RegisterUserRequest) (*model.User, error) {
	user := &model.User{
		RegisterUserRequest: *request,
	}
	return user, u.db.Create(user).Error
}

func (u *UserRepo) GetByEmail(email string) (*model.User, error) {
	user := &model.User{}
	if err := u.db.Where("email = ?", email).Limit(1).Find(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
