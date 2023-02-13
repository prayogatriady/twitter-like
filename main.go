package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/prayogatriady/twitter-like/controller"
	"github.com/prayogatriady/twitter-like/middleware"
	"github.com/prayogatriady/twitter-like/repository"
	"github.com/prayogatriady/twitter-like/utils"
)

func main() {
	db, err := utils.InitDB()
	if err != nil {
		log.Println(err)
	}

	userRepo := repository.NewUserRepo(db)
	tweetRepo := repository.NewTweetRepo(db)

	userCont := controller.NewUserCont(userRepo, tweetRepo)
	tweetCont := controller.NewTweetCont(tweetRepo)

	r := gin.Default()

	r.POST("/signup", userCont.SignUp)
	r.POST("/login", userCont.Login)
	r.Use(middleware.AuthMiddleware) // Middleware for authentication

	api := r.Group("/api")
	{
		api.GET("/profile", userCont.Profile)
		api.POST("/tweet", tweetCont.Tweet)

	}

	if err := r.Run(":8000"); err != nil {
		log.Println(err)
	}
}
