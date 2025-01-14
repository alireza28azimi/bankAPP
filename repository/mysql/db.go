package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	UserName string
	Password string
	Host     string
	Port     int
	DbName   string
}

type MySqlDB struct {
	config Config
	db     *sql.DB
}

func New(config Config) *MySqlDB {

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s port=%d",
		config.UserName, config.Password, config.DbName, config.Host, config.Port)

	db, err := sql.Open("MySqlDB ", dsn)
	if err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &MySqlDB{config: config, db: db}
}
