package icity_sdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_SetNickName(t *testing.T) {
	user := login()

	nickname := "大猪猪"
	ok := user.SetNickName(nickname)
	assert.True(t, ok)
}

func TestUser_SetSocialWechat(t *testing.T) {
	user := login()

	wechat := "hello"
	ok := user.SetSocialWechat(wechat)
	assert.True(t, ok)
}
