package icity_sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_NewDiary(t *testing.T) {
	user := login()

	title := "Test"
	comment := "Test"

	resp := user.NewDiary(title, comment, Private)
	assert.True(t, resp.Success)
	assert.NotZero(t, resp.ActivityToken)
}

func TestUser_DeleteDiary(t *testing.T) {
	user := login()

	title := "Test"
	comment := "Test"

	newResp := user.NewDiary(title, comment, Private)
	assert.True(t, newResp.Success)
	assert.NotZero(t, newResp.ActivityToken)

	deleteResp := user.DeleteDiary(newResp.ActivityToken)
	assert.True(t, deleteResp.Success)
	assert.Equal(t, deleteResp.ActivityToken, newResp.ActivityToken)
}

func TestUser_Like(t *testing.T) {
	user := login()

	title := "Test"
	comment := "Test"

	newResp := user.NewDiary(title, comment, Private)
	assert.True(t, newResp.Success)
	assert.NotZero(t, newResp.ActivityToken)

	ok := user.Like(newResp.ActivityToken)
	assert.True(t, ok)

	_ = user.DeleteDiary(newResp.ActivityToken)
}

func TestUser_UnLike(t *testing.T) {
	user := login()

	title := "Test"
	comment := "Test"

	newResp := user.NewDiary(title, comment, Private)
	assert.True(t, newResp.Success)
	assert.NotZero(t, newResp.ActivityToken)

	id := newResp.ActivityToken
	ok := user.Like(id)
	assert.True(t, ok)

	ok = user.Unlike(id)
	assert.True(t, ok)

	_ = user.DeleteDiary(newResp.ActivityToken)
}
