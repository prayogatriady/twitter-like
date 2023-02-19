package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prayogatriady/twitter-like/entities"
	"github.com/prayogatriady/twitter-like/middleware"
	"github.com/prayogatriady/twitter-like/repository"
)

type UserContInterface interface {
	SignUp(c *gin.Context)
	Login(c *gin.Context)
	// Logout(c *gin.Context)

	Profile(c *gin.Context)
	EditProfile(c *gin.Context)
}

type UserCont struct {
	userRepoInterface  repository.UserRepoInterface
	tweetRepoInterface repository.TweetRepoInterface
}

func NewUserCont(user repository.UserRepoInterface, tweet repository.TweetRepoInterface) UserContInterface {
	return &UserCont{
		userRepoInterface:  user,
		tweetRepoInterface: tweet,
	}
}

func (u *UserCont) SignUp(c *gin.Context) {
	// binding signup user json request
	var signupUser entities.SignupUser
	if err := c.ShouldBindJSON(&signupUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400 - BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	var user entities.User
	user = entities.User{
		Username:    signupUser.Username,
		Email:       signupUser.Email,
		Password:    signupUser.Password,
		ProfilePict: signupUser.ProfilePict,
	}

	user, err := u.userRepoInterface.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200 - STATUS OK",
		"message": "User created",
	})
}

func (u *UserCont) Login(c *gin.Context) {
	var userLogin entities.LoginUser
	if err := c.ShouldBindJSON(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400 - BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	user := entities.User{
		Username: userLogin.Username,
		Password: userLogin.Password,
	}

	user, err := u.userRepoInterface.GetUserByUsernamePassword(user.Username, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	// create token
	token, err := middleware.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error()},
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200 - OK",
		"message": "Login",
		"data":    token,
	})
}

func (u *UserCont) Profile(c *gin.Context) {
	// get payload from token
	userId, err := middleware.ExtractToken(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	// get user details
	user, err := u.userRepoInterface.GetUser(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	// get tweet for corresponding user
	tweets, err := u.tweetRepoInterface.GetTweets(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	// add to user struct for response
	user.Tweets = tweets

	c.JSON(http.StatusOK, gin.H{
		"status":  "200 - STATUS OK",
		"message": "Profile",
		"body":    user,
	})
}

func (u *UserCont) EditProfile(c *gin.Context) {
	// get payload from token
	userId, err := middleware.ExtractToken(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	// binding update user json request
	var updateUser entities.UpdateUser
	if err := c.ShouldBindJSON(&updateUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400 - BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	var user entities.User
	user = entities.User{
		Username:    updateUser.Username,
		Email:       updateUser.Email,
		Password:    updateUser.Password,
		ProfilePict: updateUser.ProfilePict,
	}

	user, err = u.userRepoInterface.UpdateUser(userId, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200 - OK",
		"message": "User updated",
		"body":    user,
	})
}
