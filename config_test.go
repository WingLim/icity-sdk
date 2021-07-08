package icity_sdk

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteConfig(t *testing.T) {
	filepath := "tmp.json"

	config := &Config{
		Username: "123",
		Password: "123",
	}

	err := WriteConfig(config, filepath)
	assert.Nil(t, err)

	t.Cleanup(func() {
		os.Remove(filepath)
	})
}

func TestReadConfig(t *testing.T) {
	filepath := "tmp.json"

	config := &Config{
		Username: "123",
		Password: "123",
	}

	err := WriteConfig(config, filepath)
	assert.Nil(t, err)

	newConfig, err := ReadConfig(filepath)
	assert.Nil(t, err)
	assert.Equal(t, config.Username, newConfig.Username)
	assert.Equal(t, config.Password, newConfig.Password)

	t.Cleanup(func() {
		os.Remove(filepath)
	})
}
