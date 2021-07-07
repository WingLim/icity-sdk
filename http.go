package icity_sdk

import (
	"net/http"
	"net/url"

	"github.com/WingLim/icity-sdk/constant"
)

func (user *User) get(path string) (resp *http.Response, err error) {
	fullUrl := constant.HOME + path
	return user.client.Get(fullUrl)
}

func (user *User) postForm(path string, data url.Values) (resp *http.Response, err error) {
	fullUrl := constant.HOME + path
	return user.client.PostForm(fullUrl, data)
}
