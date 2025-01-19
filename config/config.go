package config

import (
	"main.go/repository/mysql"
	"main.go/service/authservice"
)

type Config struct {
	Auth  authservice.Config `koanf:"auth"`
	Mysql mysql.Config
}
