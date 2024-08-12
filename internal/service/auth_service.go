package service

import (
	"errors"

	"github.com/RivGames/my-knowledge-base/config"
	"github.com/RivGames/my-knowledge-base/internal/model"
	"github.com/RivGames/my-knowledge-base/internal/repository"
	"github.com/RivGames/my-knowledge-base/internal/request"
	"github.com/RivGames/my-knowledge-base/pkg/errs"
	"github.com/RivGames/my-knowledge-base/pkg/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (a *AuthService) Register(c echo.Context, request *request.RegisterUserRequest) (*model.User, error) {
	if err := c.Validate(request); err != nil {
		return nil, err
	}
	user, err := a.userRepo.FetchByEmail(request.Email)
	if user != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errs.ErrEmailAlreadyTaken
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	request.Password = string(hashedPassword)

	return a.userRepo.Create(request)
}

func (a *AuthService) Login(c echo.Context, request *request.LoginUserRequest) (*model.User, error) {
	if err := c.Validate(request); err != nil {
		return nil, err
	}
	user, err := a.userRepo.FetchByEmail(request.Email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errs.ErrInvalidCredentials
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return nil, errs.ErrInvalidCredentials
	}
	user.AccessToken, _ = jwt.CreateToken(user.ID, config.Envs.JWTSecretKey)

	return user, nil
}
