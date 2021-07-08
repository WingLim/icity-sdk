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

	id, _ := s.Find(selector.DiaryId).Attr("href")
	diary.ID = strings.Split(id, "/")[2]
	diary.Nickname = s.Find(selector.DiaryNickname).Text()

	diary.Title = s.Find(selector.DiaryTitle).Text()
	diary.Content = s.Find(selector.DiaryContent).Text()
	diary.Location = s.Find(selector.DiaryLocation).Text()
	date, _ := s.Find(selector.DiaryDate).Attr("datetime")
	diary.Date, _ = time.Parse("2006-01-02T15:04:05Z", date)

	return diary
}

func (user *User) GetWorld() []Diary {
	doc, err := user.getWithDoc(path.WORLD)
	if err != nil {
		return nil
	}

	var diarys []Diary

	doc.Find(selector.WorldDiarys).Each(func(i int, s *goquery.Selection) {
		diarys = append(diarys, parseDiary(s))
	})

	return diarys
}
