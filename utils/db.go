package utils

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	// set environtment variable for setup mysql database
	DB_USER := os.Getenv("DB_USER")
	if DB_USER == "" {
		log.Println("Environment variable DB_USER must be set")
	}

	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	if DB_PASSWORD == "" {
		log.Println("Environment variable DB_PASSWORD must be set")
	}

	DB_HOST := os.Getenv("DB_HOST")
	if DB_HOST == "" {
		log.Println("Environment variable DB_HOST must be set")
	}

	DB_PORT := os.Getenv("DB_PORT")
	if DB_PORT == "" {
		log.Println("Environment variable DB_PORT must be set")
	}

	DB_NAME := os.Getenv("DB_NAME")
	if DB_NAME == "" {
		log.Println("Environment variable DB_NAME must be set")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// db.AutoMigrate(&entities.User{}, &entities.Tweet{}, &entities.Follow{})
	// db.Migrator().CreateTable(&entities.Follow{})

	return db, nil
}
