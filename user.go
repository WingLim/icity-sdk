package icity_sdk

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/WingLim/icity-sdk/constant/data"
	"github.com/WingLim/icity-sdk/constant/path"
	"github.com/WingLim/icity-sdk/constant/selector"
	"github.com/WingLim/icity-sdk/log"
)

type User struct {
	Username string
	Password string

	UserID   string
	Nickname string
	Bio      string
	Location string

	SettingsPrivacy SettingsPrivacy

	client http.Client
}

// NewUser creates a User instance.
func NewUser(username, password string) *User {
	user := &User{
		Username: username,
		Password: password,

		SettingsPrivacy: SettingsPrivacy{
			DefaultPrivacy: Public,
			InvisibleMode:  NotInvisible,
			UnWorld:        Show,
			AboutMe:        Everyone,
			Followings:     Everyone,
			Followers:      Everyone,
			AllowComment:   Everyone,
			AllowLike:      Everyone,
			AllowMessage:   Everyone,
		},
	}
	user.initClient()
	return user
}

// Follow follows one user by user id.
func (user *User) Follow(userID string) bool {
	urlPath := fmt.Sprintf(path.Follow, userID)

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

// Unfollow unfollows one user by user id.
func (user *User) Unfollow(userID string) bool {
	urlPath := fmt.Sprintf(path.Unfollow, userID)

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
	doc, err := user.getWithDoc(path.SettingsIndex)
	if err != nil {
		return err
	}

	href, _ := doc.Find(selector.UserID).Attr("href")
	hrefArr := strings.Split(href, "/")
	user.UserID = hrefArr[2]
	user.Nickname, _ = doc.Find(selector.Nickname).Attr("value")
	user.Bio = doc.Find(selector.Bio).Text()
	user.Location, _ = doc.Find(selector.Location).Attr("value")
	return nil
}

func (user *User) getUserSettingsPrivacy() error {
	doc, err := user.getWithDoc(path.Privacy)
	if err != nil {
		return err
	}

	if v, exists := doc.Find(selector.SettingsDefaultPrivacy).Attr("value"); exists {
		value, _ := strconv.Atoi(v)
		user.SettingsPrivacy.DefaultPrivacy = DiaryPrivacy(value)
	}

	if v, exists := doc.Find(selector.SettingsPrivacy).Attr("value"); exists {
		value, _ := strconv.Atoi(v)
		user.SettingsPrivacy.InvisibleMode = InvisibleType(value)
	}

	if v, exists := doc.Find(selector.SettingsUnWorld).Attr("value"); exists {
		value, _ := strconv.Atoi(v)
		user.SettingsPrivacy.UnWorld = WorldType(value)
	}

	if v, exists := doc.Find(selector.SettingsAboutMe).Attr("value"); exists {
		value, _ := strconv.Atoi(v)
		user.SettingsPrivacy.AboutMe = ViewAccess(value)
	}

	if v, exists := doc.Find(selector.SettingsFollowings).Attr("value"); exists {
		value, _ := strconv.Atoi(v)
		user.SettingsPrivacy.Followings = ViewAccess(value)
	}

	if v, exists := doc.Find(selector.SettingsFollowers).Attr("value"); exists {
		value, _ := strconv.Atoi(v)
		user.SettingsPrivacy.Followers = ViewAccess(value)
	}

	if v, exists := doc.Find(selector.SettingsAllowComment).Attr("value"); exists {
		value, _ := strconv.Atoi(v)
		user.SettingsPrivacy.AllowComment = ViewAccess(value)
	}

	if v, exists := doc.Find(selector.SettingsAllowLike).Attr("value"); exists {
		value, _ := strconv.Atoi(v)
		user.SettingsPrivacy.AllowLike = ViewAccess(value)
	}

	if v, exists := doc.Find(selector.SettingsAllowMessage).Attr("value"); exists {
		value, _ := strconv.Atoi(v)
		user.SettingsPrivacy.AllowMessage = ViewAccess(value)
	}

	return nil
}

func doGetDiaries(user *User, urlPath string) []Diary {
	doc, err := user.getWithDoc(urlPath)
	if err != nil {
		log.Error(err)
		return nil
	}

	var diaries []Diary
	doc.Find(selector.MyDiaries).Each(func(i int, s *goquery.Selection) {
		diary := parseDiary(s, false)
		diary.Nickname = user.Nickname
		diary.UserID = user.UserID
		diaries = append(diaries, diary)
	})

	return diaries
}

// GetUserDiaries gets user index page diaries.
func (user *User) GetUserDiaries(userID string) []Diary {
	urlPath := fmt.Sprintf(path.UserHome, userID)

	return doGetDiaries(user, urlPath)
}

// GetUserAllDiaries gets user all diaries.
func (user *User) GetUserAllDiaries(userID string) []Diary {
	page := 1

	var diaries []Diary
	for {
		urlPath := fmt.Sprintf(path.UserDiariesPage, userID, page)

		data := doGetDiaries(user, urlPath)
		if len(data) == 0 {
			break
		}
		diaries = append(diaries, data...)
		page++
	}
	return diaries
}

// GetMyDiaries gets my index page diaries.
func (user *User) GetMyDiaries() []Diary {
	return user.GetUserDiaries(user.UserID)
}

// GetMyAllDiaries gets my all diaries.
func (user *User) GetMyAllDiaries() []Diary {
	return user.GetUserAllDiaries(user.UserID)
}

// ChangePassword changes my password.
func (user *User) ChangePassword(currentPassword, newPassword, confirmPassword string) bool {
	postData := url.Values{}
	postData.Set(data.MethodKEY, "put")
	postData.Set(data.CurrentPassword, currentPassword)
	postData.Set(data.ResetPassword, newPassword)
	postData.Set(data.PasswordConfirmation, confirmPassword)

	resp, err := user.postForm(path.SetPassword, postData)
	if err != nil {
		log.Error(err)
		return false
	}

	if resp.StatusCode == http.StatusOK {
		return true
	}

	return false
}
