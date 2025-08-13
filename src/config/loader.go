package config

import (
	"encoding/json"
	"log"
	"os"
)

func LoadConfig() Config {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal("Cannot open file:", err)
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal("Error decoding JSON:", err)
	}

	return config
}
