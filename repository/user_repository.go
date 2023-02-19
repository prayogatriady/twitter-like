package repository

import (
	"github.com/prayogatriady/twitter-like/entities"
	"gorm.io/gorm"
)

// Interface: class-like
type UserRepoInterface interface {
	CreateUser(user entities.User) (entities.User, error)
	GetUser(userId int64) (entities.User, error)
	GetUserByUsernamePassword(username string, password string) (entities.User, error)
	UpdateUser(userId int64, updateUser entities.User) (entities.User, error)
}

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepoInterface {
	return &UserRepo{
		DB: db,
	}
}

func (r *UserRepo) CreateUser(user entities.User) (entities.User, error) {
	if err := r.DB.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepo) GetUser(userId int64) (entities.User, error) {
	var user entities.User
	if err := r.DB.Where("id = ?", userId).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepo) GetUserByUsernamePassword(username string, password string) (entities.User, error) {
	var user entities.User
	if err := r.DB.Where("username = ? AND password = ?", username, password).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepo) UpdateUser(userId int64, updateUser entities.User) (entities.User, error) {
	var user entities.User
	if err := r.DB.Where("id = ?", userId).Updates(&updateUser).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
