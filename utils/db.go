package utils

import (
	"fmt"

	"github.com/prayogatriady/twitter-like/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	USERNAME := "root"
	PASSWORD := "root"
	PORT := 3306
	DBNAME := "twitter_like"
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		USERNAME, PASSWORD, PORT, DBNAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&entities.User{}, &entities.Tweet{})

	return db, nil
}
