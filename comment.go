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
	UserID   string
	Content  string
	Date     time.Time
}

// NewComment creates a new comment of a diary by diary id.
func (user *User) NewComment(diaryID, comment string) (newResp Response) {
	urlPath := fmt.Sprintf(path.NewComment, diaryID)

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
func (user *User) DeleteComment(commentID, diaryID string) (deleteResp Response) {
	urlPath := fmt.Sprintf(path.DeleteComment, commentID, diaryID)

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
func (user *User) ReplyComment(userID, diaryID, comment string) Response {
	comment = fmt.Sprintf("@%s %s", userID, comment)

	return user.NewComment(diaryID, comment)
}

// GetComments gets diary comments by diary id.
func (user *User) GetComments(diaryID string) []Comment {
	urlPath := fmt.Sprintf(path.GetComments, diaryID)

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
