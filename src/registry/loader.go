package registry

import (
	"encoding/json"
	"log"
	"os"
)

func LoadRegistry() Registry {
	file, err := os.Open("registry.json")
	if err != nil {
		log.Fatal("Cannot open file:", err)
	}
	defer file.Close()

	var registry Registry
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&registry)
	if err != nil {
		log.Fatal("Error decoding JSON:", err)
	}

	return registry
}
