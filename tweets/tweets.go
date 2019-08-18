package tweets

import (
	"context"

	"github.com/twitter-clone/models"
)

// Repository represent the users's repository contract
type Repository interface {
	GetTweets(ctx context.Context, userid ...string) ([]models.Tweet, error)
	Insert(ctx context.Context, user *models.Tweet) error
}

type Tweets interface {
	Validate() error
}
