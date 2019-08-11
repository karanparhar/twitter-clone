package repository

import (
	"context"

	log "github.com/sirupsen/logrus"

	"time"

	"github.com/ethereum_project/config"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/ethereum_project/models"
	ethereum "github.com/ethereum_project/usecase/ethereum_repository"
)

type blocksRepository struct {
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

func (m *blocksRepository) getCollection(dbname, collection string) (c *mgo.Collection) {
	c = m.Conn.DB(dbname).C(collection)
	return c
}

// NewBlocksRepository will create an object that represent the ethereum.Repository interface
func NewBlocksRepository(Conn *mgo.Session) ethereum.Repository {
	return &blocksRepository{Conn}
}

func (m *blocksRepository) fetch(ctx context.Context, dbname, collection string, query map[string]interface{}) ([]models.Block, error) {
	var result []models.Block
	c := m.getCollection(dbname, collection)

	err := c.Find(query).All(&result)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return result, nil
}
func (m *blocksRepository) fetchone(ctx context.Context, dbname, collection string, query map[string]interface{}) (models.Block, error) {
	var result models.Block
	c := m.getCollection(dbname, collection)

	err := c.Find(query).One(&result)
	if err != nil {
		log.Error(err)
		return result, err
	}

	return result, nil
}

func (m *blocksRepository) insert(dbname, collection string, docs interface{}) error {

	c := m.getCollection(dbname, collection)

	return c.Insert(docs)

}

func (m *blocksRepository) Fetch(ctx context.Context) ([]models.Block, error) {

	res, err := m.fetch(ctx, "ethereum", "blocks", nil)
	if err != nil {
		return nil, err
	}

	return res, err
}
func (m *blocksRepository) FetchOne(ctx context.Context, query bson.M) (models.Block, error) {

	res, err := m.fetchone(ctx, "ethereum", "blocks", query)
	if err != nil {
		return res, err
	}

	return res, err
}

func (m *blocksRepository) Store(ctx context.Context, docs interface{}) error {

	return m.insert("ethereum", "blocks", docs)

}
