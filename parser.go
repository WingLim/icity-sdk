package icity_sdk

import (
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/WingLim/icity-sdk/constant/selector"
)

const timeLayout = "2006-01-02T15:04:05Z"

func parseDiary(s *goquery.Selection) Diary {
	diary := Diary{}

	id, _ := s.Find(selector.DiaryId).Attr("href")
	diary.ID = strings.Split(id, "/")[2]
	user := s.Find(selector.DiaryNickname).Text()
	nameArr := strings.Split(user, "@")
	diary.Nickname = nameArr[0]
	diary.UserId = nameArr[1]
	diary.Title = s.Find(selector.DiaryTitle).Text()
	diary.Content = s.Find(selector.DiaryContent).Text()
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
