package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// Represents database server and credentials
type Config struct {
	Server   string
	Database string
}

// Read and parse the configuration file
func (c *Config) Read() {
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
