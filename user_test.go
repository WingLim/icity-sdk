package icity_sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_Follow(t *testing.T) {
	user := login()

	userId := "winglims"

	ok := user.Follow(userId)
	assert.True(t, ok)
}

func TestUser_UnFollow(t *testing.T) {
	user := login()

	userId := "winglims"

	ok := user.UnFollow(userId)
	assert.True(t, ok)
}
