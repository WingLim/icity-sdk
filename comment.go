package icity_sdk

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/url"

	"github.com/WingLim/icity-sdk/constant/data"
	"github.com/WingLim/icity-sdk/constant/path"
)

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
		log.Println(err)
		return
	}
	return
}

func (user *User) DeleteComment(commentId string) (deleteResp Response) {
	urlPath := fmt.Sprintf(path.DELETECOMMENT, commentId)

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
