package repository

import (
	"github.com/RivGames/my-knowledge-base/internal/model"
	"github.com/RivGames/my-knowledge-base/internal/request"
	"gorm.io/gorm"
)

type UserRepository interface {
	Fetch() ([]*model.User, error)
	FetchById(uint) (*model.User, error)
	FetchByEmail(string) (*model.User, error)
	Create(*request.RegisterUserRequest) (*model.User, error)
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

func (u *UserRepo) FetchById(id uint) (*model.User, error) {
	var user *model.User
	if err := u.db.Find(user, id).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepo) FetchByEmail(email string) (*model.User, error) {
	user := &model.User{}
	if err := u.db.Where("email = ?", email).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepo) Create(req *request.RegisterUserRequest) (*model.User, error) {
	user := &model.User{
		RegisterUserRequest: *req,
	}
	return user, u.db.Create(user).Error
}
