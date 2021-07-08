package icity_sdk

import (
	"fmt"
	"net/http"

	"github.com/WingLim/icity-sdk/constant/data"
	"github.com/WingLim/icity-sdk/constant/path"
	"github.com/WingLim/icity-sdk/log"
)

type InvisibleType int

func (i InvisibleType) String() string {
	return fmt.Sprintf("%d", i)
}

const (
	NotInvisible InvisibleType = iota + 1
	ToStrangers
	ToAll
)

type ViewAccess int

func (v ViewAccess) String() string {
	return fmt.Sprintf("%d", v)
}

const (
	Everyone ViewAccess = iota + 1
	Friend
	Self
)

type WorldType int

func (w WorldType) String() string {
	return fmt.Sprintf("%d", w)
}

const (
	Show WorldType = iota
	NoShow
)

type SettingsPrivacy struct {
	DefaultPrivacy DiaryPrivacy
	InvisibleMode  InvisibleType
	UnWorld        WorldType

	AboutMe    ViewAccess
	Followings ViewAccess
	Followers  ViewAccess

	AllowComment ViewAccess
	AllowLike    ViewAccess
	AllowMessage ViewAccess
}

func (user *User) doPrivacySettings(settings ...Setting) bool {
	postData := buildSettingsData(settings)

	resp, err := user.postForm(path.PRIVACY, postData)
	if err != nil {
		log.Error(err)
		return false
	}

	if resp.StatusCode == http.StatusOK {
		return true
	}
	return false
}

func (user *User) SetDefaultPrivacy(setting DiaryPrivacy) bool {
	value := setting.String()
	set := Setting{
		Key:   data.DefaultPrivacy,
		Value: value,
	}

	ok := user.doPrivacySettings(set)
	if ok {
		user.SettingsPrivacy.DefaultPrivacy = setting
		return true
	}

	return false
}

func (user *User) SetInvisibleMode(setting InvisibleType) bool {
	value := setting.String()
	set := Setting{
		Key:   data.Privacy,
		Value: value,
	}

	ok := user.doPrivacySettings(set)
	if ok {
		user.SettingsPrivacy.InvisibleMode = setting
		return true
	}

	return false
}

func (user *User) SetUnWorld(setting WorldType) bool {
	value := setting.String()
	set := Setting{
		Key:   data.UnWorld,
		Value: value,
	}

	ok := user.doPrivacySettings(set)
	if ok {
		user.SettingsPrivacy.UnWorld = setting
		return true
	}

	return false
}

func (user *User) SetAboutMeAccess(setting ViewAccess) bool {
	value := setting.String()
	set := Setting{
		Key:   data.AboutMe,
		Value: value,
	}

	ok := user.doPrivacySettings(set)
	if ok {
		user.SettingsPrivacy.AboutMe = setting
		return true
	}

	return false
}

func (user *User) SetFollowingsAccess(setting ViewAccess) bool {
	value := setting.String()
	set := Setting{
		Key:   data.Followings,
		Value: value,
	}

	ok := user.doPrivacySettings(set)
	if ok {
		user.SettingsPrivacy.Followings = setting
		return true
	}

	return false
}

func (user *User) SetFollowersAccess(setting ViewAccess) bool {
	value := setting.String()
	set := Setting{
		Key:   data.Followers,
		Value: value,
	}

	ok := user.doPrivacySettings(set)
	if ok {
		user.SettingsPrivacy.Followers = setting
		return true
	}

	return false
}

func (user *User) SetAllowComment(setting ViewAccess) bool {
	value := setting.String()
	set := Setting{
		Key:   data.AllowComment,
		Value: value,
	}

	ok := user.doPrivacySettings(set)
	if ok {
		user.SettingsPrivacy.AllowComment = setting
		return true
	}

	return false
}

func (user *User) SetAllowLike(setting ViewAccess) bool {
	value := setting.String()
	set := Setting{
		Key:   data.AllowLike,
		Value: value,
	}

	ok := user.doPrivacySettings(set)
	if ok {
		user.SettingsPrivacy.AllowLike = setting
		return true
	}

	return false
}

func (user *User) SetAllowMessage(setting ViewAccess) bool {
	value := setting.String()
	set := Setting{
		Key:   data.AllowMessage,
		Value: value,
	}

	ok := user.doPrivacySettings(set)
	if ok {
		user.SettingsPrivacy.AllowMessage = setting
		return true
	}

	return false
}
