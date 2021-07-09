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
	assert.NotZero(t, conversation.Messages)
}

func TestUser_GetMoreMessages(t *testing.T) {
	user := login()

	conversationID := "vmo5lz5"
	conversation := user.GetConversation(conversationID)

	newConversation := user.GetMoreMessages(conversation)

	assert.NotEqual(t, len(conversation.Messages), len(newConversation.Messages))
}

func TestUser_SendMessage(t *testing.T) {
	user := login()

	conversationID := "2i94bn2"

	ok := user.SendMessage(conversationID, "hi hi~")
	assert.True(t, ok)
}
