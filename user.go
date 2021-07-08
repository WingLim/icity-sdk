package icity_sdk

import (
	"fmt"
	"net/http"

	"github.com/WingLim/icity-sdk/constant/path"
	"github.com/WingLim/icity-sdk/constant/selector"
	"github.com/WingLim/icity-sdk/log"
)

type User struct {
	Username string
	Password string

	client http.Client
}

// NewUser creates a User instance.
func NewUser(username, password string) *User {
	user := &User{
		Username: username,
		Password: password,
	}
	user.initClient()
	return user
}

// Follow follows one user by user id.
func (user *User) Follow(userId string) bool {
	urlPath := fmt.Sprintf(path.FOLLOW, userId)

	resp, err := user.post(urlPath, "", nil, iCRenderToUserHeader)
	if err != nil {
		log.Error(err)
		return false
	}

	if resp.StatusCode == http.StatusOK {
		return true
	}
	return false
}

// UnFollow unfollows one user by user id.
func (user *User) UnFollow(userId string) bool {
	urlPath := fmt.Sprintf(path.UNFOLLOW, userId)

	resp, err := user.delete(urlPath, iCRenderToUserHeader)
	if err != nil {
		log.Error(err)
		return false
	}

	if resp.StatusCode == http.StatusOK {
		return true
	}
	return false
}

func (user *User) getUserInfo() error {
	doc, err := user.getWithDoc(path.SETTINGSINDEX)
	if err != nil {
		return err
	}

	user.Nickname, _ = doc.Find(selector.NICKNAME).Attr("value")
	user.Bio = doc.Find(selector.BIO).Text()
	user.Location, _ = doc.Find(selector.LOCATION).Attr("value")
	return nil
}
