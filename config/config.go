package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"

	"os"
)

var Conf Config

type Config struct {
	DatabaseName   string   `json:"databasename"`
	User           string   `json:"user"`
	Password       string   `json:"password"`
	MongoIPs       []string `json:"mongoip"`
	Collectionname string   `json:"collection"`
}

var File = flag.String("configfile", "properties.json", "properties file location")

func GetConfig() Config {

	flag.Parse()
	file, err := os.Open(*File)
	if err != nil {
		fmt.Println("file error:", err)
		os.Exit(1)
	}
	defer file.Close()
	var configuration Config
	err = json.NewDecoder(file).Decode(&configuration)

	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	return configuration
}

// GetSession function
func GetSession(c Config) *mgo.Session {
	info := &mgo.DialInfo{
		Addrs:    c.MongoIPs,
		Timeout:  1 * time.Second,
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

func init() {
	log.Println("started")
	Conf = GetConfig()

}
