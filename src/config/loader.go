package config

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

var config Config
var once sync.Once

func LoadConfig() Config {
	once.Do(func() {
		file, err := os.Open("config.json")
		if err != nil {
			log.Fatal("Cannot open file:", err)
		}
		defer file.Close()

		decoder := json.NewDecoder(file)
		err = decoder.Decode(&config)
		if err != nil {
			log.Fatal("Error decoding JSON:", err)
		}
	})

	return config
}
