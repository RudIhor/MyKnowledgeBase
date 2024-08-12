package main

import (
	"log"

	"github.com/RivGames/my-knowledge-base/cmd/routes"
	"github.com/RivGames/my-knowledge-base/config"
	"github.com/RivGames/my-knowledge-base/internal/storage/postgresql"
	"github.com/RivGames/my-knowledge-base/pkg/app"
)

func main() {
	db, err := postgresql.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	app := app.NewApp(config.Envs, db)
	router := routes.NewRouter(app)

	router.ListenAndServe()
}
