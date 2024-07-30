package service

import (
	"errors"

	"github.com/RivGames/my-knowledge-base/config"
	"github.com/RivGames/my-knowledge-base/internal/model"
	"github.com/RivGames/my-knowledge-base/internal/repository"
	requests "github.com/RivGames/my-knowledge-base/internal/requests/auth"
	"github.com/RivGames/my-knowledge-base/pkg/errs"
	"github.com/RivGames/my-knowledge-base/pkg/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (u *UserService) Register(c echo.Context, request *requests.RegisterUserRequest) (*model.User, error) {
	if err := c.Validate(request); err != nil {
		return nil, err
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	request.Password = string(hashedPassword)
	request.AccessToken, _ = jwt.CreateToken(request.Email, config.Envs.JWTSecretKey)

	return u.userRepo.Create(request)
}

func (u *UserService) Login(c echo.Context, request *requests.LoginUserRequest) (*model.User, error) {
	if err := c.Validate(request); err != nil {
		return nil, err
	}
	user, err := u.userRepo.GetByEmail(request.Email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errs.ErrInvalidCredentials
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return nil, errs.ErrInvalidCredentials
	}
	user.AccessToken, _ = jwt.CreateToken(user.Email, config.Envs.JWTSecretKey)

	return user, nil
}
