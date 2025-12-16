package configs

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Port string
}

func SetupEnv() (cfg AppConfig, err error) {

	if os.Getenv("APP_ENV") == "dev" {
		godotenv.Load()
	}

	http_port := os.Getenv("HTTP_PORT")

	if len(http_port) < 1 {
		return AppConfig{}, errors.New("HTTP_PORT is not set")
	}

	return AppConfig{Port: http_port}, nil
}
