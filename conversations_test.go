package icity_sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_GetConversationsList(t *testing.T) {
	user := login()

	list := user.GetConversationsList()
	if len(list) != 0 {
		assert.NotZero(t, list[0].ID)
	}
}
