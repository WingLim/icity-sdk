package icity_sdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_Follow(t *testing.T) {
	user := login()

	userId := "winglims"

	ok := user.Follow(userId)
	assert.True(t, ok)
}

func TestUser_Unfollow(t *testing.T) {
	user := login()

	userId := "winglims"

	ok := user.Unfollow(userId)
	assert.True(t, ok)
}
