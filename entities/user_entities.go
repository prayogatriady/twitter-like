package entities

import "gorm.io/gorm"

// Struct for accesing database table
type User struct {
	gorm.Model
	Username string `json:"username" gorm:"size:20;unique"`
	Email    string `json:"email" gorm:"size:255;unique"`
	Password string `json:"password" gorm:"size:255"`
}

// Struct for login
type LoginUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Struct for showing profile
type ProfileUser struct {
	Username string   `json:"username"`
	Email    string   `json:"email"`
	Tweets   []string `json:"tweets"`
}
