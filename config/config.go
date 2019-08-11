package config

import (
	"encoding/json"
	"flag"
	"fmt"

	log "github.com/sirupsen/logrus"

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

func init() {
	log.Println("started")
	Conf = GetConfig()

}
