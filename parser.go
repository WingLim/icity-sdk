package icity_sdk

import (
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/WingLim/icity-sdk/constant/selector"
)

const (
	timeLayout           = "2006-01-02T15:04:05Z"
	deviceLastUsedLayout = "2006-01-02 15:04:05 -0700"
)

func parseDiary(s *goquery.Selection, world bool) Diary {
	diary := Diary{}

	id, _ := s.Find(selector.DiaryId).Attr("href")
	diary.ID = strings.Split(id, "/")[2]
	if world {
		user := s.Find(selector.DiaryNickname).Text()
		nameArr := strings.Split(user, "@")
		diary.Nickname = nameArr[0]
		diary.UserID = nameArr[1]
	}
	diary.Title = s.Find(selector.DiaryTitle).Text()
	var photos []string
	s.Find(selector.DiaryPhotos).Each(func(i int, s *goquery.Selection) {
		photos = append(photos, s.AttrOr("src", ""))
	})
	diary.Content, _ = s.Find(selector.DiaryContent).Html()
	diary.Location = s.Find(selector.DiaryLocation).Text()
	date, _ := s.Find(selector.DiaryDate).Attr("datetime")
	diary.Date, _ = time.Parse(timeLayout, date)

	return diary
}

func parseComment(s *goquery.Selection) Comment {
	comment := Comment{}

	user := s.Find(selector.DiaryNickname).Text()
	nameArr := strings.Split(user, "@")
	comment.Nickname = nameArr[0]
	comment.UserID = nameArr[1]
	comment.Content = s.Find(selector.DiaryContent).Text()
	date, _ := s.Find(selector.CommentDate).Attr("datetime")
	comment.Date, _ = time.Parse(timeLayout, date)

	return comment
}

func parseConversationItem(s *goquery.Selection) ConversationItem {
	href, _ := s.Find(selector.ConversationItemID).Attr("href")
	hrefArr := strings.Split(href, "/")

	item := ConversationItem{
		ID:          hrefArr[3],
		Nickname:    s.Find(selector.ConversationItemUser).Text(),
		LastMessage: s.Find(selector.ConversationItemLastMessage).Nodes[0].LastChild.Data,
	}

	lastDate, _ := s.Find(selector.ConversationItemLastDate).Attr("datetime")
	item.LastDate, _ = time.Parse(timeLayout, lastDate)

	return item
}

func parseConversation(s *goquery.Selection) Message {
	msg := Message{
		Content:   s.Find(selector.ConversationContent).Text(),
		Timestamp: s.AttrOr("data-sort-by", ""),
	}
	if s.HasClass("me") {
		msg.Type = TypeMe
	} else {
		msg.Type = TypeTa
	}

	return msg
}

func parseDevice(s *goquery.Selection) Device {
	device := Device{
		App:        s.Find(selector.DeviceApp).Text()[0:12],
		Hardware:   s.Find(selector.DeviceHardware).Text(),
		RemoveLink: s.Find(selector.DeviceRemoveLink).AttrOr("href", ""),
	}
	date := s.Find(selector.DeviceLastUsed).Text()
	device.LastUsed, _ = time.Parse(deviceLastUsedLayout, date)
	return device
}
