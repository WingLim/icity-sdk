package icity_sdk

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/WingLim/icity-sdk/constant/path"
	"github.com/WingLim/icity-sdk/constant/selector"
	"time"
)

type ConversationItem struct {
	ID          string
	Nickname    string
	LastMessage string
	LastDate    time.Time
}

func (user *User) GetConversationsList() []ConversationItem {
	doc, err := user.getWithDoc(path.Conversations)
	if err != nil {
		return nil
	}

	var list []ConversationItem
	doc.Find(selector.ConversationsList).Each(func(i int, s *goquery.Selection) {
		list = append(list, parseConversationItem(s))
	})

	return list
}

func (user *User) GetConversations() {

}

func (user *User) GetMoreConversations() {

}
