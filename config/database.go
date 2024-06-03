package config

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

// https://medium.com/propertyfinder-engineering/go-and-mysql-setting-up-connection-pooling-4b778ef8e560
// đây là biến quản lý connection pool, không phải là một connection
var DB *gorm.DB

type DatabaseConfig struct {
	Host     string
	Database string
	Port     int
	Username string
	Password string
}

func ConfigDatabase(config DatabaseConfig) (*gorm.DB, error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		config.Username, config.Password, config.Host, config.Port, config.Database)
	db, error := gorm.Open("mysql", dsn)
	if error != nil {
		return nil, error
	}
	sqlDb := db.DB()
	sqlDb.SetMaxOpenConns(30)
	sqlDb.SetMaxIdleConns(30)
	return db, nil
}

func InitConnection() {
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
