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

func TestUser_GetMyDiaries(t *testing.T) {
	user := login()

	diaries := user.GetMyDiaries()
	assert.NotZero(t, diaries)
}

func TestUser_GetMyAllDiaries(t *testing.T) {
	user := login()

	diaries := user.GetMyAllDiaries()
	assert.NotZero(t, diaries)
}
