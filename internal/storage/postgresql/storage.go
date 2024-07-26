package postgresql

import (
	"embed"
	"errors"
	"fmt"
	"log"

	"github.com/RivGames/my-knowledge-base/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//go:embed db/migrations/*.sql
var fs embed.FS

type PostgresStore struct {
	Db *gorm.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.Envs.DBHost, config.Envs.DBUser, config.Envs.DBPassword, config.Envs.DBName, config.Envs.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.New("Problem with connection to Postgres")
	}
	d, err := iofs.New(fs, "db/migrations")
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithSourceInstance(
		"iofs",
		d,
		fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", config.Envs.DBUser, config.Envs.DBPassword, config.Envs.DBHost, config.Envs.DBName),
	)
	if err != nil {
		return nil, err
	}
	if err := m.Up(); err != migrate.ErrNoChange {
		return nil, errors.New("Migrating failed")
	}

	return &PostgresStore{db}, nil
}
