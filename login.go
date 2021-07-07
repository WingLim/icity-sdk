package icity_sdk

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/WingLim/icity-sdk/constant"
	data2 "github.com/WingLim/icity-sdk/constant/data"
	"github.com/WingLim/icity-sdk/constant/selector"
	"log"
	"net/http"
	"net/url"
)

func (user *User) getToken() error {
	resp, err := user.get(constant.WELCOME)
	if err != nil {
		return err
	}
	defer closeBody(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return err
	}

	if token, ok := doc.Find(selector.TOKEN).Attr("value"); ok {
		user.token = token
		return nil
	}
	return errors.New("fail to get token")
}

func (user *User) buildData() url.Values {
	data := url.Values{}
	data.Set(data2.Utf8KEY, data2.DefaultUTF8)
	data.Set(data2.TokenKEY, user.token)
	data.Set(data2.UsernameKEY, user.Username)
	data.Set(data2.PasswordKEY, user.Password)
	data.Set(data2.CommitKEY, data2.DefaultCOMMIT)
	data.Set(data2.RememberKEY, data2.DefaultRemember)
	return data
}

func (user *User) checkLoginStatus() bool {
	resp, err := user.get(constant.WORLD)
	if err != nil {
		return false
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return false
	}

	text := doc.Find(selector.WORLD).Text()
	if text != "World" {
		return false
	}
	return true
}

func Login(username, password string, saveCookies bool) *User {
	user := NewUser(username, password)
	if saveCookies {
		cookies := readCookiesFromFile()
		if len(cookies) == 0 {
			goto doLogin
		}
		cookieUrl, err := url.Parse(constant.HOME)
		if err != nil {
			return nil
		}
		user.client.Jar.SetCookies(cookieUrl, cookies)
		if user.checkLoginStatus() {
			return user
		}
	}

doLogin:
	err := user.getToken()
	if err != nil {
		log.Fatal(err)
	}
	data := user.buildData()

	resp, err := user.postForm(constant.SIGNIN, data)
	if err != nil {
		log.Fatal(err)
	}
	defer closeBody(resp.Body)

	if saveCookies {
		if err = saveCookiesToFile(resp.Cookies()); err != nil {
			return nil
		}
	}

	if len(resp.Cookies()) != 0 {
		return user
	}

	return nil
}
