package postgresql

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type Config struct {
	UserName string
	Password string
	Host     string
	Port     int
	DbName   string
}

type PostgresDB struct {
	config Config
	db     *sql.DB
}

func New(config Config) *PostgresDB {

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s port=%d",
		config.UserName, config.Password, config.DbName, config.Host, config.Port)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &PostgresDB{config: config, db: db}
}
