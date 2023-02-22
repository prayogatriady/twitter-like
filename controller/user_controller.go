package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prayogatriady/twitter-like/entities"
	"github.com/prayogatriady/twitter-like/middleware"
	"github.com/prayogatriady/twitter-like/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserContInterface interface {
	SignUp(c *gin.Context)
	Login(c *gin.Context)

	HomePage(c *gin.Context)
	Profile(c *gin.Context)
	EditProfile(c *gin.Context)
}

type UserCont struct {
	userRepoInterface   repository.UserRepoInterface
	tweetRepoInterface  repository.TweetRepoInterface
	followRepoInterface repository.FollowRepoInterface
}

func NewUserCont(user repository.UserRepoInterface, tweet repository.TweetRepoInterface, follow repository.FollowRepoInterface) UserContInterface {
	return &UserCont{
		userRepoInterface:   user,
		tweetRepoInterface:  tweet,
		followRepoInterface: follow,
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

	// Generate password hash
	bytePassword, err := bcrypt.GenerateFromPassword([]byte(signupUser.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	var user entities.User
	user = entities.User{
		Username:    signupUser.Username,
		Email:       signupUser.Email,
		Password:    string(bytePassword),
		ProfilePict: signupUser.ProfilePict,
	}

	user, err = u.userRepoInterface.CreateUser(user)
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
		"body":    user,
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

	// check password
	userFound, err := u.userRepoInterface.GetUserByUsername(userLogin.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	// compare found password from database and user input password
	if err := bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(userLogin.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "401 - UNAUTHORIZED",
			"message": err.Error(),
		})
		return
	}

	user, err := u.userRepoInterface.GetUserByUsernamePassword(userLogin.Username, userFound.Password)
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

func (u *UserCont) HomePage(c *gin.Context) {
	// get payload from token
	userId, err := middleware.ExtractToken(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	follows, err := u.followRepoInterface.GetFollowing(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	var following_ids []int64
	for _, follow := range follows {
		following_ids = append(following_ids, follow.FollowingID)
	}

	tweets, err := u.tweetRepoInterface.GetTweetsByIds(following_ids)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200 - STATUS OK",
		"message": "Tweet fetched successfully",
		"body":    tweets,
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

	// Generate password hash
	bytePassword, err := bcrypt.GenerateFromPassword([]byte(updateUser.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	var user entities.User
	user = entities.User{
		Username:    updateUser.Username,
		Email:       updateUser.Email,
		Password:    string(bytePassword),
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
