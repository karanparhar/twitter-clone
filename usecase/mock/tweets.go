package mock

import (
	mock "github.com/stretchr/testify/mock"

	"github.com/twitter-clone/models"
	"github.com/twitter-clone/tweets"
)

type tweet struct {
	mock.Mock
}

func Newtweets(t models.Tweet) tweets.Tweets {

	return &tweet{}

}

func (t *tweet) Validate() error {

	return nil
}
