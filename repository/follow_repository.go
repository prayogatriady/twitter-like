package repository

import (
	"github.com/prayogatriady/twitter-like/entities"
	"gorm.io/gorm"
)

// Interface: class-like
type FollowRepoInterface interface {
	FollowUser(userIDFollower int64, userIDFollowing int64) (entities.Follow, error)
	GetFollower(userId int64) ([]entities.Follow, error)
	GetFollowing(userId int64) ([]entities.Follow, error)
}

type FollowRepo struct {
	DB *gorm.DB
}

func NewFollowRepo(db *gorm.DB) FollowRepoInterface {
	return &FollowRepo{
		DB: db,
	}
}

func (f *FollowRepo) FollowUser(userIDFollower int64, userIDFollowing int64) (entities.Follow, error) {
	var follow entities.Follow
	follow.FollowerID = userIDFollower
	follow.FollowingID = userIDFollowing

	if err := f.DB.Create(&follow).Error; err != nil {
		return follow, err
	}
	return follow, nil
}

func (f *FollowRepo) GetFollower(userId int64) ([]entities.Follow, error) {
	var follows []entities.Follow
	if err := f.DB.Where("following_id = ?", userId).Find(&follows).Error; err != nil {
		return follows, err
	}
	return follows, nil
}

func (f *FollowRepo) GetFollowing(userId int64) ([]entities.Follow, error) {
	var follows []entities.Follow
	if err := f.DB.Where("follower_id = ?", userId).Find(&follows).Error; err != nil {
		return follows, err
	}
	return follows, nil
}
