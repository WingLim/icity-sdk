package icity_sdk

import (
	"github.com/WingLim/icity-sdk/constant/path"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetLoginToken(t *testing.T) {
	user := NewUser("", "")

	token := user.getLoginToken()
	assert.NotZero(t, token)
}

func TestGetLogoutToken(t *testing.T) {
	user := Login(myUsername, myPassword, false)

	token := user.getCSRFToken()
	assert.NotZero(t, token)
}

func TestLogin(t *testing.T) {
	// This is set in config.go, if you need to test with your own account,
	// rewrite it.
	user := Login(myUsername, myPassword, true)

	resp, err := user.get(path.WORLD)
	assert.Nil(t, err)
	assert.Equal(t, resp.StatusCode, http.StatusOK)

	assert.NotNil(t, user)
}

func TestLogout(t *testing.T) {
	user := Login(myUsername, myPassword, false)

	err := Logout(user)

	assert.Nil(t, err)
}
