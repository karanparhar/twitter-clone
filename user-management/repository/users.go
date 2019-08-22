package usersrepository

import (
	"context"

	log "github.com/sirupsen/logrus"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/twitter-clone/models"
	users "github.com/twitter-clone/user-management"
)

type usersRepository struct {
	Conn *mgo.Session
}

func (m *usersRepository) getCollection(dbname, collection string) (c *mgo.Collection) {
	c = m.Conn.DB(dbname).C(collection)
	return c
}

// New will create an object that represent the twitter-clone.Repository interface

func NewUsersRepo(Conn *mgo.Session) users.Repository {
	return &usersRepository{Conn}
}

func (m *usersRepository) getuser(ctx context.Context, dbname, collection string, query map[string]interface{}) (models.Profile, error) {
	var result models.Profile
	c := m.getCollection(dbname, collection)

	err := c.Find(query).One(&result)
	if err != nil {
		log.Error(err)
		return result, err
	}

	return result, nil
}

func (m *usersRepository) insert(dbname, collection string, docs interface{}) error {

	c := m.getCollection(dbname, collection)

	return c.Insert(docs)

}
func (m *usersRepository) update(dbname, collection string, who map[string]interface{}, query map[string]interface{}) error {

	c := m.getCollection(dbname, collection)

	return c.Update(who, query)

}
func (m *usersRepository) UpdateFollowers(ctx context.Context, user models.Profile, followerid string) error {
	userdata := bson.M{"username": user.Username}
	query := bson.M{"$push": bson.M{"following": followerid}}

	err := m.update("twitter-clone", "users", userdata, query)

	return err
}

func (m *usersRepository) GetUser(ctx context.Context, user models.Profile) (*models.Profile, error) {
	query := bson.M{"username": user.Username}
	res, err := m.getuser(ctx, "twitter-clone", "users", query)
	if err != nil {
		return nil, err
	}

	return &res, err
}

func (m *usersRepository) Insert(ctx context.Context, user *models.Profile) error {

	return m.insert("twitter-clone", "users", user)

}
