package icity_sdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_GetDevices(t *testing.T) {
	user := login()

	devices := user.GetDevices()
	assert.NotZero(t, devices)
}

func TestUser_RemoveDevice(t *testing.T) {
	user := login()

	devices := user.GetDevices()
	ok := user.RemoveDevice(devices[0])
	assert.True(t, ok)
}
