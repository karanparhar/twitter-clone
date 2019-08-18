package usersrepository

import (
	"io/ioutil"
	"testing"

	"gopkg.in/mgo.v2/dbtest"
	_ "gopkg.in/tomb.v2"
)

func TestNewBlocksRepository(t *testing.T) {

	var Server dbtest.DBServer

	tempDir, _ := ioutil.TempDir("", "testing")
	Server.SetPath(tempDir)

	// My main session var is now set to the temporary MongoDB instance
	Session := Server.Session()

	NewUsersRepo(Session)

}
