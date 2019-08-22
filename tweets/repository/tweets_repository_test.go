package tweetsrepository

import (
	"context"
	"io/ioutil"
	"testing"
	"time"

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

	a := NewTweetsRepo(Session)
	tweet := models.Tweet{
		Username: "test",
		Text:     "this in my first tweet",
		Time:     time.Now().Unix(),
	}

	err := a.Insert(context.TODO(), &tweet)

	assert.NoError(t, err)

	out, err := a.GetTweets(context.TODO(), "test")
	assert.NoError(t, err)
	assert.Equal(t, tweet.Username, out[0].Username)

}
