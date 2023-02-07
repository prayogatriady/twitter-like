package entities

import "gorm.io/gorm"

type Tweet struct {
	gorm.Model
	Content  string `json:"content" gorm:"size:1000"`
	Username string `json:"username" gorm:"size:20"`
}
