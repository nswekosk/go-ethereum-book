package main

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Address   string `json:"address"`
	ProjectID string `json:"projectID"`
}

func LoadConfig(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		log.Fatalf("Error loading config: %v", err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}
