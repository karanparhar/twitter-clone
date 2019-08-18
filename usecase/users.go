package users

import (
	"github.com/twitter-clone/models"
)

// Usecase represent the  usecases
type Usecase interface {
	SignUp(user models.Profile) error
	Login(user models.Profile) (string, error)
	InsertTweets(tw models.Tweet) error
	FollowUser(u models.Profile, followerid string) error

	GetFollowersTweets(u models.Profile) ([]models.Tweet, error)
}
