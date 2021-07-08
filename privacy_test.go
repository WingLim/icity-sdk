package icity_sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_SetPrivacy(t *testing.T) {
	user := login()

	resp := user.NewDiary("", "hi", Public)
	diaryID := resp.ActivityToken

	var ok bool
	ok = user.SetOnlyFriend(diaryID)
	assert.True(t, ok)

	ok = user.SetPrivate(diaryID)
	assert.True(t, ok)

	ok = user.SetPublic(diaryID)
	assert.True(t, ok)

	t.Cleanup(func() {
		user.DeleteDiary(diaryID)
	})
}
