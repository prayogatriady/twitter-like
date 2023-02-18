package entities

import "gorm.io/gorm"

// Struct for accesing user table from database
type User struct {
	gorm.Model
	Username    string   `json:"username" gorm:"size:20;unique"`
	Email       string   `json:"email" gorm:"size:255;unique"`
	Password    string   `json:"password" gorm:"size:255"`
	ProfilePict string   `json:"profile_pic" gorm:"size:255"`
	Tweets      []Tweet  `gorm:"foreignKey:UserID;references:ID"`
	FollowerID  []Follow `gorm:"foreignKey:UserID;references:FollowerID"`
	FollowingID []Follow `gorm:"foreignKey:UserID;references:FollowingID"`
}

// Struct for login
type LoginUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Struct for showing profile
type ProfileUser struct {
	UserID   uint     `json:"user_id"`
	Username string   `json:"username"`
	Email    string   `json:"email"`
	Tweets   []string `json:"tweets"`
}
