package icity_sdk

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/WingLim/icity-sdk/constant/path"
	"github.com/WingLim/icity-sdk/constant/selector"
	"github.com/WingLim/icity-sdk/log"
)

type User struct {
	Username string
	Password string

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
			FollowingsList: Everyone,
			FollowersList:  Everyone,
			AllowComment:   Everyone,
			AllowLike:      Everyone,
			AllowMessage:   Everyone,
		},
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

func (user *User) getUserSettingsPrivacy() error {
	doc, err := user.getWithDoc(path.PRIVACY)
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
		user.SettingsPrivacy.FollowingsList = ViewAccess(value)
	}

	if v, exists := doc.Find(selector.SettingsFollowers).Attr("value"); exists {
		value, _ := strconv.Atoi(v)
		user.SettingsPrivacy.FollowersList = ViewAccess(value)
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
