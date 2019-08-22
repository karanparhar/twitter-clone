package tweet

import (
	"testing"

	"github.com/twitter-clone/models"
)

func TestNewTweet(t *testing.T) {

	tw := models.Tweet{
		Username: "test",
		Text:     "This is my first tweet",
	}
	newtweet := NewTweet(tw)

	err := newtweet.Validate()

	if err != nil {

		t.Error(err)

	}

}
