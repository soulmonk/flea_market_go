package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// todo required camelcase, do not now yet why
type PG struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
}

type Mongo struct {
	Server   string
	Database string
}

// Represents database server and credentials
type Config struct {
	Mongo Mongo
	Pg    PG
}

// read and parse the configuration file
func (c *Config) read() {
	file, e := ioutil.ReadFile("./config.json")
	if e != nil {
		log.Fatal(e)
		os.Exit(1)
	}

	var err = json.Unmarshal(file, &c)
	if err != nil {
		log.Fatal("Cannot unmarshal the json ", err)
	}
}

func Load() *Config {
	config := Config{}
	config.read()

	return &config
}
