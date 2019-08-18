package tweetsrepository

import (
	"context"

	log "github.com/sirupsen/logrus"

	"time"

	"github.com/twitter-clone/config"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/twitter-clone/models"
	tweets "github.com/twitter-clone/tweets"
)

type tweetsRepository struct {
	Conn *mgo.Session
}

// GetSession function
func GetSession(c config.Config) *mgo.Session {
	info := &mgo.DialInfo{
		Addrs:    c.MongoIPs,
		Timeout:  60 * time.Second,
		Database: c.DatabaseName,
		Username: c.User,
		Password: c.Password,
	}

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		log.Fatalf("ERROR: Not Able to Connect to MongoDB")
	}

	session.SetMode(mgo.Monotonic, true)

	return session
}

func (m *tweetsRepository) getCollection(dbname, collection string) (c *mgo.Collection) {
	c = m.Conn.DB(dbname).C(collection)
	return c
}

// New will create an object that represent the twitter-clone.Repository interface

func NewTweetsRepo(Conn *mgo.Session) tweets.Repository {
	return &tweetsRepository{Conn}
}

func (m *tweetsRepository) gettweets(ctx context.Context, dbname, collection string, query map[string]interface{}) ([]models.Tweet, error) {
	var result []models.Tweet
	c := m.getCollection(dbname, collection)

	err := c.Find(query).Sort("-time").All(&result)
	if err != nil {
		log.Error(err)
		return result, err
	}

	return result, nil
}

func (m *tweetsRepository) insert(dbname, collection string, docs interface{}) error {

	c := m.getCollection(dbname, collection)

	return c.Insert(docs)

}
func (m *tweetsRepository) GetTweets(ctx context.Context, usersid ...string) ([]models.Tweet, error) {
	query := bson.M{"username": bson.M{"$in": usersid}}
	res, err := m.gettweets(ctx, "twitter-clone", "tweets", query)
	if err != nil {
		return nil, err
	}

	return res, err
}

func (m *tweetsRepository) Insert(ctx context.Context, tweet *models.Tweet) error {

	return m.insert("twitter-clone", "tweets", tweet)

}
