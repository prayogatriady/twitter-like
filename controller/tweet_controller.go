package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prayogatriady/twitter-like/entities"
	"github.com/prayogatriady/twitter-like/middleware"
	"github.com/prayogatriady/twitter-like/repository"
)

type TweetContInterface interface {
	Tweet(c *gin.Context)
}

type TweetCont struct {
	TweetRepoInterface repository.TweetRepoInterface
}

func NewTweetCont(TweetRepoInterface repository.TweetRepoInterface) TweetContInterface {
	return &TweetCont{
		TweetRepoInterface: TweetRepoInterface,
	}
}

func (t *TweetCont) Tweet(c *gin.Context) {
	// get data from token
	userId, err := middleware.ExtractToken(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	var postTweet entities.PostTweet
	if err := c.ShouldBindJSON(&postTweet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400 - BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	tweet := entities.Tweet{
		Content: postTweet.Content,
		UserID:  userId,
	}

	tweet, err = t.TweetRepoInterface.CreateTweet(tweet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200 - STATUS OK",
		"message": "Tweet Posted",
		"body":    tweet,
	})
}
