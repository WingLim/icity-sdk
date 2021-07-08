package icity_sdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_NewComment(t *testing.T) {
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

func TestUser_ReplyComment(t *testing.T) {
	user := login()

	commentId := "iwqbwn7"
	userId := "winglims"
	comment := "I am replying you!"

	resp := user.ReplyComment(userId, commentId, comment)
	assert.True(t, resp.Success)
	assert.NotZero(t, resp.ActivityToken)
}

func TestUser_GetComments(t *testing.T) {
	user := login()

	id := "73sxpl4"

	comments := user.GetComments(id)
	assert.NotZero(t, comments)
}
