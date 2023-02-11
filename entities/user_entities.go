package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"size:20;unique"`
	Email    string `json:"email" gorm:"size:255;unique"`
	Password string `json:"password" gorm:"size:255"`
}

type LoginUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
