package app

import (
	"net/http"

	"github.com/RivGames/my-knowledge-base/internal/storage/postgresql"
	"github.com/labstack/echo/v4"
)

type Server struct {
	echo  *echo.Echo
	store *postgresql.PostgresStore
}

func NewServer(store *postgresql.PostgresStore) *Server {
	return &Server{
		echo:  echo.New(),
		store: store,
	}
}

func (s *Server) Run() {
	s.echo.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello My Knowledge Base!")
	})
	s.echo.Logger.Fatal(s.echo.Start(":80"))
}
