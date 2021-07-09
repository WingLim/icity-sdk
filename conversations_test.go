package icity_sdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_GetConversationsList(t *testing.T) {
	user := login()

	list := user.GetConversationsList()
	if len(list) != 0 {
		assert.NotZero(t, list[0].ID)
	}
}

func TestUser_GetConversation(t *testing.T) {
	user := login()

	list := user.GetConversationsList()

	conversation := user.GetConversation(list[0].ID)
	assert.NotZero(t, conversation)
}

func TestUser_SendMessage(t *testing.T) {
	user := login()

	conversationID := "2i94bn2"

	ok := user.SendMessage(conversationID, "hi hi~")
	assert.True(t, ok)
}