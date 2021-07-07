package icity_sdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewComment(t *testing.T) {
	user := login()

	diaryId := "ou4yza9"
	comment := "hey boy!"
	resp := user.NewComment(diaryId, comment)
	assert.True(t, resp.Success)
	assert.NotZero(t, resp.ActivityToken)
}

func TestUser_DeleteComment(t *testing.T) {
	user := login()

	diaryId := "ou4yza9"
	comment := "hey boy!"
	newResp := user.NewComment(diaryId, comment)
	assert.True(t, newResp.Success)
	assert.NotZero(t, newResp.ActivityToken)

	id := newResp.ActivityToken
	deleteResp := user.DeleteComment(id)
	assert.True(t, deleteResp.Success)
	assert.Equal(t, id, deleteResp.ActivityToken)
}
