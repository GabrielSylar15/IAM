package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

// https://medium.com/propertyfinder-engineering/go-and-mysql-setting-up-connection-pooling-4b778ef8e560
var DB *sql.DB

type DatabaseConfig struct {
	Host     string
	Database string
	Port     int
	Username string
	Password string
}

func ConfigDatabase(config DatabaseConfig) (*sql.DB, error) {
	dsn := fmt.Sprintln("%s:%s@tcp(%s:%d)/%s&readTimeout=30s&writeTimeout=30s",
		config.Username, config.Password, config.Host, config.Port, config.Database)
	db, error := sql.Open("mysql", dsn)
	if error != nil {
		return nil, error
	}
	db.SetMaxOpenConns(30)
	db.SetMaxIdleConns(30)
	return db, nil
}

func init() {
	config := DatabaseConfig{
		Username: os.Getenv("DATABASE_MYSQL_USERNAME"),
		Password: os.Getenv("DATABASE_MYSQL_PASSWORD"),
		Host:     os.Getenv("DATABASE_MYSQL_HOST"),
		Port:     3306,
		Database: os.Getenv("DATABASE_MYSQL_SCHEME"),
	}

	var err error
	DB, err = ConfigDatabase(config)
	if err != nil {
		log.Fatalf("Could not set up database connection: %v", err)
	}
}
