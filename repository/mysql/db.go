package mysql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	UserName string
	Password string
	Port     int
	Host     string
	DbName   string
}

type MysqlDB struct {
	config Config
	db     *sql.DB
}

// GetUserByPhoneNumber implements uservalidator.Repository.

func New(config Config) *MysqlDB {
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s", config.UserName, config.Password, config.Host, config.Port, config.DbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &MysqlDB{config: config, db: db}
}
