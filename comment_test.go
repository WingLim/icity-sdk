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
