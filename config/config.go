package config

import (
	"main.go/repository/mysql"
)

type Config struct {
	Mysql mysql.Config
}
