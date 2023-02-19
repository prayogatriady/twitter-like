package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prayogatriady/twitter-like/entities"
	"github.com/prayogatriady/twitter-like/middleware"
	"github.com/prayogatriady/twitter-like/repository"
)

type FollowContInterface interface {
	FollowUser(c *gin.Context)
	GetFollower(c *gin.Context)
	GetFollowing(c *gin.Context)
}

type FollowCont struct {
	FollowRepoInterface repository.FollowRepoInterface
}

func NewFollowCont(FollowRepoInterface repository.FollowRepoInterface) FollowContInterface {
	return &FollowCont{
		FollowRepoInterface: FollowRepoInterface,
	}
}

func (f *FollowCont) FollowUser(c *gin.Context) {
	// get payload from token
	userId, err := middleware.ExtractToken(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	var postFollow entities.PostFollow
	if err := c.ShouldBindJSON(&postFollow); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400 - BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	follow, err := f.FollowRepoInterface.FollowUser(userId, postFollow.FollowingID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200 - STATUS OK",
		"message": "Follow user successfully",
		"body":    follow,
	})
}

func (f *FollowCont) GetFollower(c *gin.Context) {
	// get payload from token
	userId, err := middleware.ExtractToken(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	follows, err := f.FollowRepoInterface.GetFollower(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200 - STATUS OK",
		"message": "Follower Found",
		"body":    follows,
	})
}

func (f *FollowCont) GetFollowing(c *gin.Context) {
	// get payload from token
	userId, err := middleware.ExtractToken(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	follows, err := f.FollowRepoInterface.GetFollowing(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200 - STATUS OK",
		"message": "Following Found",
		"body":    follows,
	})
}
