package icity_sdk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/WingLim/icity-sdk/constant/data"
	"github.com/WingLim/icity-sdk/constant/path"
	"github.com/WingLim/icity-sdk/log"
)

type Diary struct {
	ID       string
	Nickname string
	UserId   string

	Title    string
	Content  string
	Location string
	Date     time.Time
}

// generateHeaders generates some headers for make sure we can have
// correct response from iCity.
func generateHeaders(user *User) []Header {
	return []Header{
		csrfHeader(user.getCSRFToken()),
		refererHeader(path.HOME + "/"),
		xRequestedWithHeader,
		acceptHeader,
	}
}

// NewDiary creates a new diary with title, content and privacy.
func (user *User) NewDiary(title, content string, privacy DiaryPrivacy) (newResp Response) {
	postData := url.Values{}
	postData.Set(data.TitleKEY, title)
	postData.Set(data.CommentKEY, content)
	postData.Set(data.PrivacyKEY, privacy.String())

	headers := generateHeaders(user)
	resp, err := user.postForm(path.NEWDIARY, postData, headers...)
	if err != nil {
		log.Error(err)
		return
	}
	defer closeBody(resp.Body)

	res, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(res, &newResp)
	if err != nil {
		log.Error(err)
		return
	}
	return
}

// DeleteDiary deletes the diary with given diary id.
func (user *User) DeleteDiary(diaryId string) (deleteResp Response) {
	urlPath := fmt.Sprintf(path.DELETEDIARY, diaryId)

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

// Like likes a diary with given diary id.
func (user *User) Like(diaryId string) bool {
	urlPath := fmt.Sprintf(path.LIKE, diaryId)

	resp, err := user.post(urlPath, "", nil, iCRenderToSelfHeader)
	if err != nil {
		log.Error(err)
		return false
	}
	defer closeBody(resp.Body)

	if resp.StatusCode == http.StatusOK {
		return true
	}
	return false
}

// UnLike unlikes a diary with given diary id.
func (user *User) UnLike(diaryId string) bool {
	urlPath := fmt.Sprintf(path.UNLIKE, diaryId)

	resp, err := user.delete(urlPath, iCRenderToSelfHeader)
	if err != nil {
		log.Error(err)
		return false
	}
	defer closeBody(resp.Body)

	if resp.StatusCode == http.StatusOK {
		return true
	}
	return false
}
