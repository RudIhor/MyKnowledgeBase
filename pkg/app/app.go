package app

import (
	"github.com/RivGames/my-knowledge-base/config"
	"github.com/RivGames/my-knowledge-base/internal/storage/postgresql"
)

type App struct {
	Config config.Config
	Store  *postgresql.PostgresStore
}

func NewApp(config config.Config, db *postgresql.PostgresStore) *App {
	return &App{config, db}
}
