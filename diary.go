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

type Privacy int

func (p Privacy) String() string {
	return fmt.Sprintf("%d", p)
}

const (
	Public Privacy = iota + 1
	OnlyFriend
	Private
)

type Response struct {
	Success       bool   `json:"success"`
	ActivityToken string `json:"activity_token"`
}

func generateHeaders(user *User) []Header {
	return []Header{
		csrfHeader(user.getCSRFToken()),
		refererHeader(path.HOME + "/"),
		xRequestedWithHeader,
		acceptHeader,
	}
}

func (user *User) NewDiary(title, comment string, privacy Privacy) (newResp Response) {
	postData := url.Values{}
	postData.Set(data.TitleKEY, title)
	postData.Set(data.CommentKEY, comment)
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
