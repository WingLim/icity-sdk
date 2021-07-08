package icity_sdk

import (
	"fmt"
	"log"
	"net/http"

	"github.com/WingLim/icity-sdk/constant/path"
)

type User struct {
	Username string
	Password string

	client http.Client
}

func NewUser(username, password string) *User {
	user := &User{
		Username: username,
		Password: password,
	}
	user.initClient()
	return user
}

func (user *User) Follow(userId string) bool {
	urlPath := fmt.Sprintf(path.FOLLOW, userId)

	resp, err := user.post(urlPath, "", nil, iCRenderToUserHeader)
	if err != nil {
		log.Println(err)
		return false
	}

	if resp.StatusCode == http.StatusOK {
		return true
	}
	return false
}

func (user *User) UnFollow(userId string) bool {
	urlPath := fmt.Sprintf(path.UNFOLLOW, userId)

	resp, err := user.delete(urlPath, iCRenderToUserHeader)
	if err != nil {
		log.Println(err)
		return false
	}

	if resp.StatusCode == http.StatusOK {
		return true
	}
	return false
}
