package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

type (
	Config struct {
		GRPC
		Log
		Postgres
	}

	GRPC struct {
		Port string `env:"GRPC_PORT,required"`
	}

	Log struct {
		Level string `env:"LOG_LEVEL,required"`
	}

	Postgres struct {
		MaxOpenConns int    `env:"DB_MAX_OPEN_CONNS"`
		MaxIdleConns int    `env:"DB_MAX_IDLE_CONNS"`
		Host         string `env:"DB_HOST,required"`
		Port         int    `env:"DB_PORT,required"`
		User         string `env:"DB_USER,required"`
		Password     string `env:"DB_PASSWORD,required"`
		DbName       string `env:"DB_NAME,required"`
		SSLMode      string `env:"DB_SSL_MODE,required"`
	}
)

func NewConfig() (*Config, error) {
	cfg := Config{}

	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return &cfg, nil
}
