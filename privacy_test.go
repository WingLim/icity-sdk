package icity_sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetPrivacy(t *testing.T) {
	user := login()

	resp := user.NewDiary("", "hi", Public)
	id := resp.ActivityToken

	var ok bool
	ok = user.SetOnlyFriend(id)
	assert.True(t, ok)

	ok = user.SetPrivate(id)
	assert.True(t, ok)

	ok = user.SetPublic(id)
	assert.True(t, ok)

	user.DeleteDiary(id)
}
