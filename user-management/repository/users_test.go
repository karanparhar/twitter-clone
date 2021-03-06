package usersrepository

import (
	"context"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/twitter-clone/models"
	"gopkg.in/mgo.v2/dbtest"
	_ "gopkg.in/tomb.v2"
)

func TestNewUsersRepo(t *testing.T) {

	var Server dbtest.DBServer

	tempDir, _ := ioutil.TempDir("", "testing")
	Server.SetPath(tempDir)

	// My main session var is now set to the temporary MongoDB instance
	Session := Server.Session()

	a := NewUsersRepo(Session)
	user := models.Profile{
		Username: "test",
		Password: "testing",
	}

	err := a.Insert(context.TODO(), &user)

	assert.NoError(t, err)

	out, err := a.GetUser(context.TODO(), user)
	assert.NoError(t, err)
	assert.NotNil(t, out, nil)
	assert.Equal(t, user.Username, out.Username)

	err = a.UpdateFollowers(context.TODO(), user, "test1")

	assert.NoError(t, err)

}
