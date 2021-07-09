package icity_sdk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/WingLim/icity-sdk/constant/data"
	"github.com/WingLim/icity-sdk/constant/path"
	"github.com/WingLim/icity-sdk/constant/selector"
	"github.com/WingLim/icity-sdk/log"
)

// ConversationItem defines every conversation in the conversation list.
type ConversationItem struct {
	ID          string
	Nickname    string
	LastMessage string
	LastDate    time.Time
}

// MessageType defines who send the message.
type MessageType string

const (
	// TypeMe means message send by myself.
	TypeMe MessageType = "Me"
	// TypeTa means message send by others.
	TypeTa MessageType = "Ta"
)

// Conversation defines one conversation and the messages.
type Conversation struct {
	// Messages includes the message in the conversation.
	Messages []Message
	// MoreMessages is a link to get more message in the conversation.
	MoreMessages string
}

// Message defines the message in the conversation.
type Message struct {
	Type      MessageType
	Content   string
	Timestamp string
}

// GetConversationsList gets all conversations we have started.
func (user *User) GetConversationsList() []ConversationItem {
	doc, err := user.getWithDoc(path.Conversations)
	if err != nil {
		log.Error(err)
		return nil
	}

	var list []ConversationItem
	doc.Find(selector.ConversationsList).Each(func(i int, s *goquery.Selection) {
		list = append(list, parseConversationItem(s))
	})

	return list
}

// GetConversation gets one conversation details.
func (user *User) GetConversation(conversationID string) *Conversation {
	urlPath := fmt.Sprintf(path.Conversation, conversationID)
	doc, err := user.getWithDoc(urlPath)
	if err != nil {
		log.Error(err)
		return nil
	}

	conversation := &Conversation{
		MoreMessages: doc.Find(selector.MoreMessages).AttrOr("href", ""),
	}
	doc.Find(selector.Conversations).Each(func(i int, s *goquery.Selection) {
		conversation.Messages = append(conversation.Messages, parseConversation(s))
	})

	return conversation
}

// GetMoreMessages gets more messages in the conversation.
func (user *User) GetMoreMessages(conversation *Conversation) *Conversation {
	urlPath := conversation.MoreMessages
	if urlPath == "" {
		return conversation
	}

	doc, err := user.getWithDoc(urlPath)
	if err != nil {
		log.Error(err)
		return conversation
	}

	conversation.MoreMessages = doc.Find(selector.MoreMessages).AttrOr("href", "")
	var moreMessages []Message
	doc.Find(selector.Conversations).Each(func(i int, s *goquery.Selection) {
		moreMessages = append(moreMessages, parseConversation(s))
	})
	moreMessages = append(moreMessages, conversation.Messages...)
	conversation.Messages = moreMessages
	return conversation
}

// SendMessage sends a message in the conversation by conversation id.
func (user *User) SendMessage(conversationID, content string) bool {
	urlPath := fmt.Sprintf(path.SendMessage, conversationID)
	postData := url.Values{}
	postData.Set(data.Content, content)

	headers := generateHeaders(user)
	resp, err := user.postForm(urlPath, postData, headers...)
	if err != nil {
		log.Error(err)
		return false
	}
	defer closeBody(resp.Body)

	res, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return false
	}

	var statusResp Response
	err = json.Unmarshal(res, &statusResp)
	if statusResp.Success {
		return true
	}

	return false
}
