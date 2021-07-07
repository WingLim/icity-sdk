package icity_sdk

import (
	"github.com/WingLim/icity-sdk/constant"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetToken(t *testing.T) {
	user := NewUser("", "")

	err := user.getToken()
	assert.Nil(t, err)
	assert.NotZero(t, user.token)
}

func TestLogin(t *testing.T) {
	// This is set in config.go, if you need to test with your own account,
	// rewrite it.
	user := Login(myUsername, myPassword, true)

	resp, err := user.get(constant.WORLD)
	assert.Nil(t, err)
	assert.Equal(t, resp.StatusCode, http.StatusOK)

	assert.NotNil(t, user)
}
