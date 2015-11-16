package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/olebedev/config"
)

const configfile = "config.json"

func LoadConfig() (cfg *config.Config) {

	file, err := os.Open(configfile)
	if err != nil {
		log.Fatalf("Cannot open file %s: %s", configfile, err)
	}

	defer file.Close()

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Cannot read contents of %s: %s", configfile, err)
	}

	cfg, err = config.ParseJson(string(buf))
	if err != nil {
		log.Fatalf("Error when parsing %s as JSON: %s", configfile, err)
	}

	return cfg
}
