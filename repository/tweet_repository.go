package repository

import (
	"github.com/prayogatriady/twitter-like/entities"
	"gorm.io/gorm"
)

type TweetRepoInterface interface {
	CreateTweet(tweet entities.Tweet) (entities.Tweet, error)
	GetTweets(username string) ([]entities.Tweet, error)
}

type TweetRepo struct {
	DB *gorm.DB
}

func NewTweetRepo(db *gorm.DB) TweetRepoInterface {
	return &TweetRepo{
		DB: db,
	}
}

func (t *TweetRepo) CreateTweet(tweet entities.Tweet) (entities.Tweet, error) {
	if err := t.DB.Create(&tweet).Error; err != nil {
		return tweet, err
	}
	return tweet, nil
}

func (t *TweetRepo) GetTweets(username string) ([]entities.Tweet, error) {
	var tweets []entities.Tweet
	if err := t.DB.Where("username = ?", username).Find(&tweets).Error; err != nil {
		return tweets, err
	}
	return tweets, nil
}
