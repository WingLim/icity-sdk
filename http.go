package icity_sdk

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/WingLim/icity-sdk/constant/path"
)

type Header struct {
	Key   string
	Value string
}

const userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36"

func (user *User) do(req *http.Request, headers ...Header) (resp *http.Response, err error) {
	req.Header.Set("User-Agent", userAgent)
	for _, h := range headers {
		req.Header.Set(h.Key, h.Value)
	}
	return user.client.Do(req)
}

func (user *User) get(urlPath string, headers ...Header) (resp *http.Response, err error) {
	fullUrl := path.HOME + urlPath
	req, err := http.NewRequest(http.MethodGet, fullUrl, nil)
	if err != nil {
		return
	}
	return user.do(req, headers...)
}

func (user *User) post(urlPath, contentType string, body io.Reader, headers ...Header) (resp *http.Response, err error) {
	fullUrl := path.HOME + urlPath
	req, err := http.NewRequest(http.MethodPost, fullUrl, body)
	if err != nil {
		return
	}
	return user.do(req, headers...)
}

func (user *User) postForm(urlPath string, data url.Values, headers ...Header) (resp *http.Response, err error) {
	return user.post(urlPath, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()), headers...)
}

func (user *User) getWithDoc(urlPath string) (*goquery.Document, error) {
	resp, err := user.get(urlPath)
	if err != nil {
		return nil, err
	}
	defer closeBody(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	return doc, nil
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
