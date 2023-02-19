package entities

import "gorm.io/gorm"

// Struct for accesing user table from database
type User struct {
	gorm.Model
	Username    string   `json:"username" gorm:"size:20;unique"`
	Email       string   `json:"email" gorm:"size:255;unique"`
	Password    string   `json:"password" gorm:"size:255"`
	ProfilePict string   `json:"profile_pict" gorm:"size:255"`
	Tweets      []Tweet  `gorm:"foreignKey:UserID;references:ID"`
	FollowerID  []Follow `gorm:"foreignKey:FollowerID;references:ID"`
	FollowingID []Follow `gorm:"foreignKey:FollowingID;references:ID"`
}

// Struct for login
type LoginUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Struct for signup
type SignupUser struct {
	Username    string `json:"username" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
	ProfilePict string `json:"profile_pict" binding:"required"`
}

// Struct for update
type UpdateUser struct {
	Username    string `json:"username"`
	Email       string `json:"email" binding:"email"`
	Password    string `json:"password"`
	ProfilePict string `json:"profile_pict"`
}

// Struct for showing profile
type ProfileUser struct {
	UserID   uint     `json:"user_id"`
	Username string   `json:"username"`
	Email    string   `json:"email"`
	Tweets   []string `json:"tweets"`
}
