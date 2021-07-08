package icity_sdk

import (
	"fmt"
	"net/http"

	"github.com/WingLim/icity-sdk/constant/path"
	"github.com/WingLim/icity-sdk/log"
)

type DiaryPrivacy int

func (p DiaryPrivacy) String() string {
	return fmt.Sprintf("%d", p)
}

const (
	Public DiaryPrivacy = iota + 1
	OnlyFriend
	Private
)

func (user *User) setPrivacy(urlPath string) bool {
	headers := generateHeaders(user)
	resp, err := user.post(urlPath, "", nil, headers...)
	if err != nil {
		log.Error(err)
		return false
	}
	if resp.StatusCode == http.StatusFound {
		return true
	}
	return false
}

// SetPublic sets diary privacy to Public.
func (user *User) SetPublic(id string) bool {
	urlPath := fmt.Sprintf(path.SETPRIVACY, id, Public)
	return user.setPrivacy(urlPath)
}

// SetOnlyFriend sets diary privacy to OnlyFriend.
func (user *User) SetOnlyFriend(id string) bool {
	urlPath := fmt.Sprintf(path.SETPRIVACY, id, OnlyFriend)
	return user.setPrivacy(urlPath)
}

// SetPrivate sets diary privacy to Private.
func (user *User) SetPrivate(id string) bool {
	urlPath := fmt.Sprintf(path.SETPRIVACY, id, Private)
	return user.setPrivacy(urlPath)
}
