package icity_sdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_GetFriends(t *testing.T) {
	user := login()

	diaries := user.GetFriends()
	assert.Len(t, diaries, 25)
}
