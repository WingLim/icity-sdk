package icity_sdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
	diaryID := newResp.ActivityToken
	assert.True(t, newResp.Success)
	assert.NotZero(t, diaryID)

	ok := user.Like(diaryID)
	assert.True(t, ok)

	t.Cleanup(func() {
		user.DeleteDiary(diaryID)
	})
}

func TestUser_UnLike(t *testing.T) {
	user := login()

	title := "Test"
	comment := "Test"

	newResp := user.NewDiary(title, comment, Private)
	assert.True(t, newResp.Success)
	assert.NotZero(t, newResp.ActivityToken)

	diaryID := newResp.ActivityToken
	ok := user.Like(diaryID)
	assert.True(t, ok)

	ok = user.Unlike(diaryID)
	assert.True(t, ok)

	t.Cleanup(func() {
		user.DeleteDiary(diaryID)
	})
}
