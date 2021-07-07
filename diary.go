package icity_sdk

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/WingLim/icity-sdk/constant/data"
	"github.com/WingLim/icity-sdk/constant/path"
)

type Response struct {
	// Success is operation status.
	Success bool `json:"success"`

	// ActivityToken is the id of diary.
	ActivityToken string `json:"activity_token"`
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
func (user *User) NewDiary(title, content string, privacy Privacy) (newResp Response) {
	postData := url.Values{}
	postData.Set(data.TitleKEY, title)
	postData.Set(data.CommentKEY, content)
	postData.Set(data.PrivacyKEY, privacy.String())

	headers := generateHeaders(user)
	resp, err := user.postForm(path.NEWDIARY, postData, headers...)
	if err != nil {
		log.Println(err)
		return
	}
	defer closeBody(resp.Body)

	res, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(res, &newResp)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

// DeleteDiary deletes the diary with given id.
func (user *User) DeleteDiary(id string) (deleteResp Response) {
	urlPath := fmt.Sprintf(path.DELETEDIARY, id)

	headers := generateHeaders(user)
	resp, err := user.delete(urlPath, headers...)
	if err != nil {
		log.Println(err)
		return
	}
	defer closeBody(resp.Body)

	res, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(res, &deleteResp)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

// Like likes a diary with given id.
func (user *User) Like(id string) bool {
	urlPath := fmt.Sprintf(path.LIKE, id)

	resp, err := user.post(urlPath, "", nil, iCRenderToHeader)
	if err != nil {
		log.Println(err)
		return false
	}
	defer closeBody(resp.Body)

	if resp.StatusCode == http.StatusOK {
		return true
	}
	return false
}

// UnLike unlikes a diary with given id.
func (user *User) UnLike(id string) bool {
	urlPath := fmt.Sprintf(path.UNLIKE, id)

	resp, err := user.delete(urlPath, iCRenderToHeader)
	if err != nil {
		log.Println(err)
		return false
	}
	defer closeBody(resp.Body)

	if resp.StatusCode == http.StatusOK {
		return true
	}
	return false
}
