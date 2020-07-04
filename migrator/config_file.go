package main

import (
	"encoding/json"
	"log"
	"os"
)

type ConfigFile struct {
	Environments map[string]map[string]interface{} `json:"environments"`
}

func NewConfigFile(path string) *ConfigFile {
	fstream, e := os.Open(path)
	if e != nil {
		log.Fatal(e)
	}

	var configFile ConfigFile
	e = json.NewDecoder(fstream).Decode(&configFile)
	if e != nil {
		log.Fatal(e)
	}

	return &configFile
}
