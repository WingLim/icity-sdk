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

type Comment struct {
	Nickname string
	UserId   string
	Content  string
	Date     time.Time
}

// NewComment creates a new comment of a diary by diary id.
func (user *User) NewComment(diaryId, comment string) (newResp Response) {
	urlPath := fmt.Sprintf(path.NEWCOMMENT, diaryId)

	postData := url.Values{}
	postData.Set(data.CommentKEY, comment)

	headers := generateHeaders(user)

	resp, err := user.postForm(urlPath, postData, headers...)
	if err != nil {
		return
	}
	defer closeBody(resp.Body)

	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &newResp)
	if err != nil {
		log.Error(err)
		return
	}
	return
}

// DeleteComment deletes a comment by comment id.
func (user *User) DeleteComment(commentId, diaryId string) (deleteResp Response) {
	urlPath := fmt.Sprintf(path.DELETECOMMENT, commentId, diaryId)

	headers := generateHeaders(user)
	resp, err := user.delete(urlPath, headers...)
	if err != nil {
		log.Error(err)
		return
	}
	defer closeBody(resp.Body)

	res, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(res, &deleteResp)
	if err != nil {
		log.Error(err)
		return
	}
	return
}

// ReplyComment replies user by user id and diary id.
func (user *User) ReplyComment(userId, diaryId, comment string) Response {
	comment = fmt.Sprintf("@%s %s", userId, comment)

	return user.NewComment(diaryId, comment)
}

// GetComments gets diary comments by diary id.
func (user *User) GetComments(diaryId string) []Comment {
	urlPath := fmt.Sprintf(path.GETCOMMENTS, diaryId)

	doc, err := user.getWithDoc(urlPath, iCRenderToRepliesHeader)
	if err != nil {
		log.Error(err)
		return nil
	}

	var comments []Comment

	doc.Find(selector.Comments).Each(func(i int, s *goquery.Selection) {
		comments = append(comments, parseComment(s))
	})

	return comments
}
