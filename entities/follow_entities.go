package entities

import (
	"gorm.io/gorm"
)

// Struct for accesing database table
type Follow struct {
	gorm.Model
	FollowerID  int64 `json:"follower_id" gorm:"index;not null"`
	FollowingID int64 `json:"following_id" gorm:"index;not null"`
}

// Struct for user follow another user
type PostFollow struct {
	FollowingID int64 `json:"following_id" binding:"required"`
}
