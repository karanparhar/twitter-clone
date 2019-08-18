package tweet

import (
	"errors"

	"github.com/twitter-clone/models"
	"github.com/twitter-clone/tweets"
)

type TweetsByTime []models.Tweet

func (t TweetsByTime) Len() int {
	return len(t)
}
func (t TweetsByTime) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
func (t TweetsByTime) Less(i, j int) bool {
	return t[i].Time > t[j].Time
}

type tweet struct {
	models.Tweet
}

func NewTweet(t models.Tweet) tweets.Tweets {

	return &tweet{t}

}

func (t *tweet) Validate() error {

	if t.Username == "" {

		return errors.New("please provide the username")
	}

	if t.Text == "" {
		return errors.New("text is empty")
	}

	return nil

}
