package icity_sdk

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/WingLim/icity-sdk/constant/path"
	"github.com/WingLim/icity-sdk/constant/selector"
	"github.com/WingLim/icity-sdk/log"
)

func (user *User) GetWorld() []Diary {
	doc, err := user.getWithDoc(path.WORLD)
	if err != nil {
		log.Error(err)
		return nil
	}

	var diaries []Diary

	doc.Find(selector.WorldDiarys).Each(func(i int, s *goquery.Selection) {
		diaries = append(diaries, parseDiary(s))
	})

	return diaries
}
