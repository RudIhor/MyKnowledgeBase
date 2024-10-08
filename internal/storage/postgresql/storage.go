package postgresql

import (
	"embed"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/RivGames/my-knowledge-base/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//go:embed db/migrations/*.sql
var fs embed.FS

type PostgresStore struct {
	DB *gorm.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.Envs.DBHost, config.Envs.DBUser, config.Envs.DBPassword, config.Envs.DBName, config.Envs.DBPort)
	logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
		},
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger,
	})
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
		return nil, err
	}

	return &PostgresStore{db}, nil
}
