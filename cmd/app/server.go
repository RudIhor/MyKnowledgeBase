package app

import (
	"net/http"

	"github.com/RivGames/my-knowledge-base/internal/repository"
	requests "github.com/RivGames/my-knowledge-base/internal/requests/auth"
	"github.com/RivGames/my-knowledge-base/internal/service"
	"github.com/RivGames/my-knowledge-base/internal/storage/postgresql"
	"github.com/RivGames/my-knowledge-base/pkg/errs"
	"github.com/RivGames/my-knowledge-base/pkg/validation"
	"github.com/labstack/echo/v4"
)

type Server struct {
	echo  *echo.Echo
	store *postgresql.PostgresStore
}

func NewServer(store *postgresql.PostgresStore) *Server {
	echo := echo.New()
	echo.Validator = validation.NewCustomValidator()
	return &Server{
		echo:  echo,
		store: store,
	}
}

func (s *Server) Run() {
	s.echo.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello My Knowledge Base!")
	})

	// Auth
	s.echo.POST("/api/v1/register", s.handleRegister)
	s.echo.POST("/api/v1/login", s.handleLogin)

	// Questions
	s.echo.GET("/api/v1/questions", s.handleGetQuestions)
	s.echo.POST("/api/v1/questions", s.handleCreateQuestion)
	s.echo.GET("/api/v1/questions/{id}", s.handleShowQuestion)
	s.echo.PUT("/api/v1/questions/{id}", s.handleUpdateQuestion)
	s.echo.DELETE("/api/v1/questions/{id}", s.handleDeleteQuestion)

	// Answers
	s.echo.GET("/api/v1/answers", s.handleGetAnswers)
	s.echo.POST("/api/v1/answers", s.handleCreateAnswer)
	s.echo.GET("/api/v1/answers/{id}", s.handleShowAnswer)
	s.echo.GET("/api/v1/question/{id}/answers", s.handleGetQuestionAnswers)
	s.echo.PUT("/api/v1/answers/{id}", s.handleUpdateAnswer)
	s.echo.DELETE("/api/v1/answers/{id}", s.handleDeleteAnswer)

	s.echo.Logger.Fatal(s.echo.Start(":80"))
}

// Questions
func (s *Server) handleGetQuestions(c echo.Context) (err error) {
	return
}

func (s *Server) handleCreateQuestion(c echo.Context) (err error) {
	return
}

func (s *Server) handleShowQuestion(c echo.Context) (err error) {
	return
}

func (s *Server) handleUpdateQuestion(c echo.Context) (err error) {
	return
}

func (s *Server) handleDeleteQuestion(c echo.Context) (err error) {
	return
}

// Answers
func (s *Server) handleGetAnswers(c echo.Context) (err error) {
	return
}

func (s *Server) handleCreateAnswer(c echo.Context) (err error) {
	return
}

func (s *Server) handleShowAnswer(c echo.Context) (err error) {
	return
}

func (s *Server) handleGetQuestionAnswers(c echo.Context) (err error) {
	return
}

func (s *Server) handleUpdateAnswer(c echo.Context) (err error) {
	return
}

func (s *Server) handleDeleteAnswer(c echo.Context) (err error) {
	return
}

// Auth
func (s *Server) handleRegister(c echo.Context) (err error) {
	registerRequest := new(requests.RegisterUserRequest)
	if err := c.Bind(registerRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	authService := service.NewAuthService(repository.NewUserRepository(s.store.DB))
	user, err := authService.Register(c, registerRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errs.ErrSomethingWentWrong)
	}

	return c.JSON(http.StatusCreated, user)
}

func (s *Server) handleLogin(c echo.Context) (err error) {
	request := new(requests.LoginUserRequest)
	if err := c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	userService := service.NewAuthService(repository.NewUserRepository(s.store.DB))
	user, err := userService.Login(c, request)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	return c.JSON(http.StatusOK, user)
}
