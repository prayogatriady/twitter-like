package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/prayogatriady/twitter-like/controller"
	"github.com/prayogatriady/twitter-like/repository"
	"github.com/prayogatriady/twitter-like/utils"
)

func main() {
	db, err := utils.InitDB()
	if err != nil {
		log.Println(err)
	}

	userRepo := repository.NewUserRepo(db)
	userCont := controller.NewUserCont(userRepo)

	r := gin.New()
	r.Use(gin.Logger())

	r.POST("/signup", userCont.SignUp)

	api := r.Group("/api")
	{
		api.GET("/profile/:username", userCont.Profile)
	}

	if err := r.Run(":8000"); err != nil {
		log.Println(err)
	}
}
