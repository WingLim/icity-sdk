package icity_sdk

import (
	"encoding/json"
	"os"
)

type Config struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// ReadConfig reads config from file.
func ReadConfig(filepath string) (config Config, err error) {
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

// WriteConfig writes config to file.
func WriteConfig(config *Config, filepath string) (err error) {
	f, err := os.Create(filepath)
	if err != nil {
		return
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	err = encoder.Encode(config)
	if err != nil {
		return
	}

	return nil
}
