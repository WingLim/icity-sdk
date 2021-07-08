package icity_sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadConfig(t *testing.T) {
	filepath := "config.json"

	config, err := readConfig(filepath)
	assert.Nil(t, err)
	assert.NotZero(t, config.Username)
	assert.NotZero(t, config.Password)
}
