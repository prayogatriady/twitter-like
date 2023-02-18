package entities

import (
	"time"

	"gorm.io/gorm"
)

// Struct for accesing database table
type Follow struct {
	FollowerID  int64 `json:"follower_id"`
	FollowingID int64 `json:"following_id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
