package icity_sdk

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/WingLim/icity-sdk/constant/path"
	"github.com/WingLim/icity-sdk/constant/selector"
	"strings"
	"time"
)

func parseDiary(s *goquery.Selection) Diary {
	diary := Diary{}

	id, _ := s.Find("div.meta > a.timeago").Attr("href")
	diary.ID = strings.Split(id, "/")[2]
	diary.Nickname = s.Find("div.line > a").Text()

	diary.Title = s.Find("div.line > h4 > a").Text()
	diary.Content = s.Find("div.line > div.comment").Text()
	diary.Location = s.Find("div.line > span.location").Text()
	date, _ := s.Find("div.meta > a.timeago > time").Attr("datetime")
	diary.Date, _ = time.Parse("2006-01-02T15:04:05Z", date)

	return diary
}

func (user *User) GetWorld() []Diary {
	doc, err := user.getWithDoc(path.WORLD)
	if err != nil {
		return nil
	}

	var diarys []Diary

	doc.Find(selector.WORLDDIARYS).Each(func(i int, s *goquery.Selection) {
		diarys = append(diarys, parseDiary(s))
	})

	return diarys
}
