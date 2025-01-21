package main

import (
	"time"

	"main.go/config"
	"main.go/delivery/httpserver"
	"main.go/repository/mysql"
	"main.go/service/authservice"
	"main.go/service/userservice"
	"main.go/service/uservalidator"
)

const (
	JwtSignKey                     = "jwt_secret"
	AccessTokenSubject             = "ac"
	RefreshTokenSubject            = "rt"
	AccessTokenExpirationDuration  = time.Hour * 24
	RefreshTokenExpirationDuration = time.Hour * 24 * 7
)

func main() {
	cfg := config.Config{
		Auth: authservice.Config{
			SignKey:               JwtSignKey,
			AccessExpirationTime:  AccessTokenExpirationDuration,
			RefreshExpirationTime: RefreshTokenExpirationDuration,
			AccessSubject:         AccessTokenSubject,
			RefreshSubject:        RefreshTokenSubject,
		},
		Mysql: mysql.Config{
			UserName: "gameapp_db",
			Password: "gameapp123",
			Port:     3308,
			Host:     "localhost",
			DbName:   "gameapp_db",
		},
	}
	authSvc, userSvc, userValidator := setupServices(cfg)
	server := httpserver.New(cfg, authSvc, userSvc, userValidator)
	server.Serve()
}
func setupServices(cfg config.Config) (authservice.Service, userservice.Service, uservalidator.Validator) {
	authSvc := authservice.New(cfg.Auth)
	MysqlRepo := mysql.New(cfg.Mysql)
	userSvc := userservice.New(authSvc, MysqlRepo)
	uv := uservalidator.New(MysqlRepo)

	return authSvc, userSvc, uv
}
