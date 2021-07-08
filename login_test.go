package icity_sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func login() *User {
	return LoginWithConfig("config.json", WithSaveCookies("cookies.json"))
}

func TestUser_GetLoginToken(t *testing.T) {
	user := NewUser("", "")

	token := user.getLoginToken()
	assert.NotZero(t, token)
}

func TestUser_GetCSRFToken(t *testing.T) {
	user := login()

	token := user.getCSRFToken()
	assert.NotZero(t, token)
}

func TestLogin(t *testing.T) {
	config, err := ReadConfig("config.json")
	assert.Nil(t, err)

	user := Login(config.Username, config.Password)

	assert.NotNil(t, user)
	assert.NotZero(t, user.UserID)
}

func TestLoginWithConfig(t *testing.T) {
	filepath := "config.json"

	user := LoginWithConfig(filepath)
	assert.NotNil(t, user)
	assert.NotZero(t, user.UserID)
}

func TestLogout(t *testing.T) {
	user := login()

	err := Logout(user)

	assert.Nil(t, err)
}
