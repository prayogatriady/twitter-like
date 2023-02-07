package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prayogatriady/twitter-like/entities"
	"github.com/prayogatriady/twitter-like/repository"
)

type UserContInterface interface {
	SignUp(c *gin.Context)
	Login(c *gin.Context)
	// Logout(c *gin.Context)

	Tweet(c *gin.Context)
	Profile(c *gin.Context)
}

type UserCont struct {
	userRepoInterface repository.UserRepoInterface
}

func NewUserCont(userRepoInterface repository.UserRepoInterface) UserContInterface {
	return &UserCont{
		userRepoInterface: userRepoInterface,
	}
}

func (u *UserCont) SignUp(c *gin.Context) {
	var user entities.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400 - BAD REQUEST",
			"message": err.Error(),
		})
		return
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
}
func (u *UserCont) Tweet(c *gin.Context) {
}
func (u *UserCont) Profile(c *gin.Context) {

	user, err := u.userRepoInterface.GetUser(c.Param("username"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200 - STATUS OK",
		"message": "Profile",
		"body":    user,
	})
}
