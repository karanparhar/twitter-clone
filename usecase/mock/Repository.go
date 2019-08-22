package mock

import (
	"context"

	mock "github.com/stretchr/testify/mock"
	"github.com/twitter-clone/models"
	"github.com/twitter-clone/tweets"
	users "github.com/twitter-clone/user-management"
)

type usersRepository struct {
	mock.Mock
}

func NewuserRepository() users.Repository {
	return &usersRepository{}

}

func NewtweetRepo() tweets.Repository {
	return &tweetRepository{}

}

type tweetRepository struct {
	mock.Mock
}

func (t *tweetRepository) GetTweets(ctx context.Context, userid ...string) ([]models.Tweet, error) {

	return nil, nil
}

func (t *tweetRepository) Insert(ctx context.Context, user *models.Tweet) error {
	return nil
}

func (r *usersRepository) Insert(ctx context.Context, user *models.Profile) error {

	return nil

}

func (r *usersRepository) GetUser(ctx context.Context, user models.Profile) (*models.Profile, error) {

	return &user, nil
}

func (r *usersRepository) UpdateFollowers(ctx context.Context, user models.Profile, followerid string) error {
	return nil
}
