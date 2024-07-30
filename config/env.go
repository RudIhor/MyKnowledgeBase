package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort      string
	DBHost       string
	DBUser       string
	DBName       string
	DBPassword   string
	DBPort       string
	TTL          string
	JWTSecretKey string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		AppPort:      getEnv("APP_PORT", "80"),
		DBHost:       getEnv("DB_HOST", "db"),
		DBUser:       getEnv("DB_USER", "postgres"),
		DBName:       getEnv("DB_NAME", "postgres"),
		DBPassword:   getEnv("DB_PASSWORD", "password"),
		DBPort:       getEnv("DB_PORT", "5432"),
		TTL:          getEnv("TTL", "15"),
		JWTSecretKey: getEnv("JWT_SECRET", "secretkey"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
