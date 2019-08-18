package users

import (
	"context"

	"github.com/twitter-clone/models"
)

// Repository represent the users's repository contract
type Repository interface {
	GetUser(ctx context.Context, user models.Profile) (*models.Profile, error)
	Insert(ctx context.Context, user *models.Profile) error
	UpdateFollowers(ctx context.Context, user models.Profile, followerid string) error
}

// Users data

type Profile interface {
	Authenticate(ctx context.Context, password string) error
	GenerateUserToken(ctx context.Context) (string, error)
	Encrypt(ctx context.Context, password string) (*models.Profile, error)
	Validate(ctx context.Context) error
}
