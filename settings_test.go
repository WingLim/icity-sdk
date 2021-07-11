package icity_sdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_SetNickName(t *testing.T) {
	user := login()

	nickname := "大猪猪"
	ok := user.SetMyNickName(nickname)
	assert.True(t, ok)
}

func TestUser_SetSocialWechat(t *testing.T) {
	user := login()

	wechat := "hello"
	ok := user.SetSocialWechat(wechat)
	assert.True(t, ok)
}

func TestUser_SetDefaultPrivacy(t *testing.T) {
	user := login()

	ok := user.SetDefaultPrivacy(Private)
	assert.True(t, ok)
	assert.Equal(t, Private, user.SettingsPrivacy.DefaultPrivacy)
}

func TestUser_SetUnWorld(t *testing.T) {
	user := login()

	ok := user.SetUnWorld(Show)
	assert.True(t, ok)
	assert.Equal(t, Show, user.SettingsPrivacy.UnWorld)
}

func TestUser_SetAboutMeAccess(t *testing.T) {
	user := login()

	ok := user.SetAboutMeAccess(Friend)
	assert.True(t, ok)
	assert.Equal(t, Friend, user.SettingsPrivacy.AboutMe)
}
