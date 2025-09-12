package app

import (
	postgres "to-do-app/pkg"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type (
	Config struct {
		Postgres postgres.PostgresConfig
	}
)

func New() (*Config, error) {
	var cfg Config
	opts := env.Options{Environment: nil, TagName: "env", Prefix: ""}

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	err = env.ParseWithOptions(&cfg, opts)
	if err != nil {
		return nil, err
	}

	return &cfg, err
}
