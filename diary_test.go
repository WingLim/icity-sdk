package icity_sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDiary(t *testing.T) {
	user := Login(myUsername, myPassword, true)

	title := "Test"
	comment := "Test"

	resp := user.NewDiary(title, comment, Private)
	assert.True(t, resp.Success)
	assert.NotZero(t, resp.ActivityToken)
}

func TestDeleteDiary(t *testing.T) {
	user := Login(myUsername, myPassword, true)

	title := "Test"
	comment := "Test"

	newResp := user.NewDiary(title, comment, Private)
	assert.True(t, newResp.Success)
	assert.NotZero(t, newResp.ActivityToken)

	deleteResp := user.DeleteDiary(newResp.ActivityToken)
	assert.True(t, deleteResp.Success)
	assert.Equal(t, deleteResp.ActivityToken, newResp.ActivityToken)
}

func TestLike(t *testing.T) {
	user := Login(myUsername, myPassword, true)

	title := "Test"
	comment := "Test"

	newResp := user.NewDiary(title, comment, Private)
	assert.True(t, newResp.Success)
	assert.NotZero(t, newResp.ActivityToken)

	ok := user.Like(newResp.ActivityToken)
	assert.True(t, ok)

	_ = user.DeleteDiary(newResp.ActivityToken)
}

func TestUnLike(t *testing.T) {
	user := Login(myUsername, myPassword, true)

	title := "Test"
	comment := "Test"

	newResp := user.NewDiary(title, comment, Private)
	assert.True(t, newResp.Success)
	assert.NotZero(t, newResp.ActivityToken)

	id := newResp.ActivityToken
	ok := user.Like(id)
	assert.True(t, ok)

	ok = user.UnLike(id)
	assert.True(t, ok)

	_ = user.DeleteDiary(newResp.ActivityToken)
}
