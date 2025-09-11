package app

import postgres "to-do-app/pkg"

type (
	Config struct {
		Postgres postgres.PostgresConfig
	}
)
