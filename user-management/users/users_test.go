package user

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/twitter-clone/models"
)

func TestNewProfile(t *testing.T) {

	u := models.Profile{
		Username: "testing",
		Password: "test123",
	}
	newprofile := NewProfile(u)

	if newprofile == nil {

		t.Failed()

	}

	user, err := newprofile.Encrypt(context.TODO(), "testing")

	if err != nil {

		t.Error(err)

	}

	assert.NotNil(t, user)

	err = newprofile.Authenticate(context.TODO(), "testing")

	if err != nil {

		t.Error(err)
	}

	err = newprofile.Validate(context.TODO())

	if err != nil {

		t.Error(err)
	}

	_, err = newprofile.GenerateUserToken(context.TODO())

	if err != nil {
		t.Error(err)
	}

}
