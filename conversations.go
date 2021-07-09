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

type Conversation struct {
	Messages     []Message
	MoreMessages string
}

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
