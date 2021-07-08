package icity_sdk

import (
	"encoding/json"
	"os"
)

type Config struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func readConfig(filepath string) (config Config, err error) {
	f, err := os.Open(filepath)
	if err != nil {
		return
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	err = decoder.Decode(&config)
	if err != nil {
		return
	}

	return config, nil
}
