package usecase

import (
	"context"
	"time"

	"github.com/twitter-clone/models"
	tweets "github.com/twitter-clone/tweets"
	tweet "github.com/twitter-clone/tweets/tweet"
	u "github.com/twitter-clone/usecase"
	usersrepository "github.com/twitter-clone/user-management"
	user "github.com/twitter-clone/user-management/users"
)

type usecase struct {
	userrepo       usersrepository.Repository
	tweetsrepo     tweets.Repository
	tweets         tweets.Tweets
	users          usersrepository.Profile
	contextTimeout context.Context
}

func NewtransactionUsecase(ctx context.Context, r usersrepository.Repository, t tweets.Repository) u.Usecase {

	return &usecase{
		userrepo:       r,
		contextTimeout: ctx,
		tweetsrepo:     t,
	}

}

func (t *usecase) NewProfile(u models.Profile) usersrepository.Profile {

	return user.NewProfile(u)

}

func (t *usecase) NewTweet(tw models.Tweet) tweets.Tweets {

	return tweet.NewTweet(tw)
}
func (t *usecase) SignUp(u models.Profile) error {

	t.users = t.NewProfile(u)

	err := t.users.Validate(t.contextTimeout)

	if err != nil {
		return err
	}

	userdata, err := t.users.Encrypt(t.contextTimeout, u.Password)

	if err != nil {
		return err
	}

	return t.userrepo.Insert(t.contextTimeout, userdata)

}
func (t *usecase) Login(u models.Profile) (string, error) {

	userdata, err := t.userrepo.GetUser(t.contextTimeout, u)

	if err != nil {
		return "", err
	}

	t.users = t.NewProfile(*userdata)
	err = t.users.Authenticate(t.contextTimeout, u.Password)

	if err != nil {
		return "", err
	}

	return t.users.GenerateUserToken(t.contextTimeout)

}

func (t *usecase) FollowUser(u models.Profile, followerid string) error {

	userdata, err := t.userrepo.GetUser(t.contextTimeout, u)

	if err != nil {
		return err
	}

	t.users = t.NewProfile(*userdata)

	return t.userrepo.UpdateFollowers(t.contextTimeout, *userdata, followerid)

}

func (t *usecase) InsertTweets(tw models.Tweet) error {

	t.tweets = t.NewTweet(tw)

	err := t.tweets.Validate()

	if err != nil {
		return err
	}

	tw.Time = time.Now().UnixNano()

	return t.tweetsrepo.Insert(t.contextTimeout, &tw)

}

func (t *usecase) GetFollowersTweets(u models.Profile) ([]models.Tweet, error) {
	userdata, err := t.userrepo.GetUser(t.contextTimeout, u)

	if err != nil {
		return nil, err
	}

	return t.tweetsrepo.GetTweets(t.contextTimeout, userdata.Following...)

}
