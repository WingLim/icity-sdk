package icity_sdk

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/WingLim/icity-sdk/constant/data"
	"github.com/WingLim/icity-sdk/constant/path"
	"github.com/WingLim/icity-sdk/constant/selector"
	"github.com/WingLim/icity-sdk/log"
)

func (user *User) getLoginToken() string {
	doc, err := user.getWithDoc(path.Welcome)
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
		log.Error(err)
		return ""
	}

	if token, ok := doc.Find(selector.CSRFToken).Attr("content"); ok {
		return token
	}
	return ""
}

func buildLoginData(username, password, token string) url.Values {
	postData := url.Values{}
	postData.Set(data.Utf8KEY, data.DefaultUtf8)
	postData.Set(data.TokenKEY, token)
	postData.Set(data.UsernameKEY, username)
	postData.Set(data.PasswordKEY, password)
	postData.Set(data.CommitKEY, "登入")
	postData.Add(data.RememberKEY, "0")
	postData.Add(data.RememberKEY, data.DefaultRemember)
	return postData
}

func buildLogoutData(token string) url.Values {
	postData := url.Values{}
	postData.Set(data.MethodKEY, "delete")
	postData.Set(data.TokenKEY, token)
	return postData
}

// checkLoginStatus checks if we have the correct permissions.
func (user *User) checkLoginStatus() bool {
	doc, err := user.getWithDoc(path.World)
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
		cookieUrl, _ := url.Parse(path.Home)
		user.client.Jar.SetCookies(cookieUrl, cookies)
		// If the cookies is expired, then login again.
		if user.checkLoginStatus() {
			return user
		}
	}

doLogin:
	token := user.getLoginToken()
	postData := buildLoginData(user.Username, user.Password, token)

	resp, err := user.postForm(path.Login, postData)
	if err != nil {
		log.Error(err)
		return nil
	}
	defer closeBody(resp.Body)

	if saveCookies {
		cookieUrl, _ := url.Parse(path.Home)
		cookies := user.client.Jar.Cookies(cookieUrl)
		if err = saveCookiesToFile(cookies); err != nil {
			log.Error(err)
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
	postData := buildLogoutData(token)

	resp, err := user.postForm(path.Logout, postData, refererHeader(path.Home+"/"))
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
	user = user.login(saveCookies)
	if user == nil {
		log.Infof("fail to log in")
		os.Exit(1)
	}
	err := user.getUserInfo()
	if err != nil {
		log.Error(err)
		return nil
	}
	err = user.getUserSettingsPrivacy()
	if err != nil {
		log.Error(err)
		return nil
	}
	return user
}

// Logout logouts user from iCity
func Logout(user *User) error {
	err := user.logout()
	if err != nil {
		return err
	}
	return nil
}
