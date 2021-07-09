package icity_sdk

import (
	"fmt"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/WingLim/icity-sdk/constant/path"
	"github.com/WingLim/icity-sdk/constant/selector"
	"github.com/WingLim/icity-sdk/log"
)

type ConversationItem struct {
	ID          string
	Nickname    string
	LastMessage string
	LastDate    time.Time
}

type MessageType string

const (
	TypeMe MessageType = "Me"
	TypeTa MessageType = "Ta"
)

type Message struct {
	Type      MessageType
	Content   string
	Timestamp string
}

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

func (user *User) GetConversation(conversationID string) []Message {
	urlPath := fmt.Sprintf(path.Conversation, conversationID)
	doc, err := user.getWithDoc(urlPath)
	if err != nil {
		log.Error(err)
		return nil
	}

	var conversation []Message
	doc.Find(selector.Conversations).Each(func(i int, s *goquery.Selection) {
		conversation = append(conversation, parseConversation(s))
	})

	return conversation
}

func (user *User) GetMoreConversations() {

}
