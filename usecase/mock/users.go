package mock

import (
	"context"

	mock "github.com/stretchr/testify/mock"
	"github.com/twitter-clone/models"
	users "github.com/twitter-clone/user-management"
)

func NewProfile() users.Profile {
	return &usersProfile{}

}

type usersProfile struct {
	mock.Mock
}

func (p *usersProfile) Encrypt(ctx context.Context, password string) (*models.Profile, error) {
	profile := &models.Profile{}
	return profile, nil

}

func (p *usersProfile) Validate(ctx context.Context) error {

	return nil
}
func (p *usersProfile) Authenticate(ctx context.Context, password string) error {

	return nil
}

func (p *usersProfile) GenerateUserToken(ctx context.Context) (string, error) {

	return "", nil
}
