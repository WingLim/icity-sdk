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

func (user *User) initClient() {
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal(err)
	}
	user.client.Jar = jar
	user.client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
}
