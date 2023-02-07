package repository

import "github.com/prayogatriady/twitter-like/entities"

type TweetRepoInterface interface {
	CreateTweet(tweet entities.Tweet) error
}
