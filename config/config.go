package config

import "main.go/repository/postgresql"

type Config struct {
	PostgresDB postgresql.Config
}
