package main

import (
	"log"

	"github.com/RivGames/my-knowledge-base/cmd/app"
	"github.com/RivGames/my-knowledge-base/internal/storage/postgresql"
)

func main() {
	db, err := postgresql.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	server := app.NewServer(db)

	server.Run()
}
