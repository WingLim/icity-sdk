package icity_sdk

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
	"github.com/WingLim/icity-sdk/constant/data"
	"github.com/WingLim/icity-sdk/constant/path"
	"github.com/WingLim/icity-sdk/constant/selector"
)

func (user *User) getLoginToken() string {
	doc, err := user.getWithDoc(path.WELCOME)
	if err != nil {
		return ""
	}

	if token, ok := doc.Find(selector.LoginToken).Attr("value"); ok {
		return token
	}
	return ""
}

func (user *User) getCSRFToken() string {
	doc, err := user.getWithDoc("/")
	if err != nil {
		return ""
	}

	if token, ok := doc.Find(selector.CSRFToken).Attr("content"); ok {
		return token
	}
	return ""
}

func (user *User) buildLoginData(token string) url.Values {
	postData := url.Values{}
	postData.Set(data.Utf8KEY, data.DefaultUtf8)
	postData.Set(data.TokenKEY, token)
	postData.Set(data.UsernameKEY, user.Username)
	postData.Set(data.PasswordKEY, user.Password)
	postData.Set(data.CommitKEY, data.DefaultCommit)
	postData.Add(data.RememberKEY, "0")
	postData.Add(data.RememberKEY, data.DefaultRemember)
	return postData
}

func (user *User) buildLogoutData(token string) url.Values {
	postData := url.Values{}
	postData.Set(data.MethodKEY, data.DefaultMethod)
	postData.Set(data.TokenKEY, token)
	return postData
}

// checkLoginStatus checks if we have the correct permissions.
func (user *User) checkLoginStatus() bool {
	resp, err := user.get(path.WORLD)
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

func (user *User) login(saveCookies bool) *User {
	if saveCookies {
		cookies := readCookiesFromFile()
		if len(cookies) == 0 {
			goto doLogin
		}
		cookieUrl, _ := url.Parse(path.HOME)
		user.client.Jar.SetCookies(cookieUrl, cookies)
		// If the cookies is expired, then login again.
		if user.checkLoginStatus() {
			return user
		}
	}

doLogin:
	token := user.getLoginToken()
	postData := user.buildLoginData(token)

	resp, err := user.postForm(path.SIGNIN, postData)
	if err != nil {
		log.Fatal(err)
	}
	defer closeBody(resp.Body)

	if saveCookies {
		cookieUrl, _ := url.Parse(path.HOME)
		cookies := user.client.Jar.Cookies(cookieUrl)
		if err = saveCookiesToFile(cookies); err != nil {
			return nil
		}
	}

	if len(resp.Cookies()) != 0 && resp.StatusCode == http.StatusFound {
		return user
	}

	return nil
}

func (user *User) logout() error {
	token := user.getCSRFToken()
	postData := user.buildLogoutData(token)

	resp, err := user.postForm(path.SIGNOUT, postData, refererHeader(path.HOME+"/"))
	if err != nil {
		return err
	}
	defer closeBody(resp.Body)

	if resp.StatusCode != http.StatusFound {
		return fmt.Errorf("fail to logout: %d %s", resp.StatusCode, resp.Status)
	}
	return nil
}

// Login logins user to iCity.
// If set true to saveCookies, then will write cookies to cookies.json,
// then will login to iCity with exists cookies.
func Login(username, password string, saveCookies bool) *User {
	user := NewUser(username, password)
	return user.login(saveCookies)
}

// Logout logouts user from iCity
func Logout(user *User) error {
	err := user.logout()
	if err != nil {
		return err
	}
	return nil
}
