package entities

import "gorm.io/gorm"

type Tweet struct {
	gorm.Model
	Content      string `json:"content" gorm:"size:1000;not null"`
	UserID       int64  `json:"user_id" gorm:"not null"`
	LikeCount    int64  `json:"like_count" gorm:"default:0"`
	RetweetCount int64  `json:"retweet_count" gorm:"default:0"`
}

type PostTweet struct {
	Content string `json:"content" binding:"required"`
}
