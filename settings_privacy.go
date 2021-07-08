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
	switch i {
	case NotInvisible:
		return "NotInvisible"
	case ToStrangers:
		return "ToStrangers"
	case ToAll:
		return "ToAll"
	default:
		return "Unknown"
	}
}

func (i InvisibleType) Data() string {
	return fmt.Sprintf("%d", i)
}

const (
	// NotInvisible shows all your diaries to everyone.
	NotInvisible InvisibleType = iota + 1

	// ToStrangers shows your diaries to your friends.
	ToStrangers

	// ToAll show your diaries to yourself.
	ToAll
)

type ViewAccess int

func (v ViewAccess) String() string {
	switch v {
	case Everyone:
		return "Everyone"
	case Friend:
		return "Friend"
	case Self:
		return "Self"
	default:
		return "Unknown"
	}
}

func (v ViewAccess) Data() string {
	return fmt.Sprintf("%d", v)
}

const (
	// Everyone means everyone on iCity have access.
	Everyone ViewAccess = iota + 1

	// Friend means only your friend have access.
	Friend

	// Self means only yourself have access.
	Self
)

type WorldType int

func (w WorldType) String() string {
	switch w {
	case Show:
		return "Shown"
	case NotShow:
		return "NotShow"
	default:
		return "Unknown"
	}
}

func (w WorldType) Data() string {
	return fmt.Sprintf("%d", w)
}

const (
	// Show publishes your diary to world.
	Show WorldType = iota

	// NotShow doesn't publish your diary to world.
	NotShow
)

type SettingsPrivacy struct {
	// DefaultPrivacy sets your new diary default access.
	DefaultPrivacy DiaryPrivacy

	// InvisibleMode sets invisible mode.
	InvisibleMode InvisibleType

	// UnWorld sets if your diary publish to world
	UnWorld WorldType

	// AboutMe sets who can access your about me page.
	AboutMe ViewAccess

	// Followings sets who can access your followings list.
	Followings ViewAccess

	// Followers sets who can access your followers list.
	Followers ViewAccess

	// AllowComment sets who can comment your diary.
	AllowComment ViewAccess

	// Allow Like sets who can like your diary.
	AllowLike ViewAccess

	// AllowMessage sets who can send you private message.
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

// SetDefaultPrivacy sets DefaultPrivacy access.
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

// SetInvisibleMode sets InvisibleMode.
func (user *User) SetInvisibleMode(setting InvisibleType) bool {
	value := setting.Data()
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

// SetUnWorld sets UnWorld type.
func (user *User) SetUnWorld(setting WorldType) bool {
	value := setting.Data()
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

// SetAboutMeAccess sets AboutMe access.
func (user *User) SetAboutMeAccess(setting ViewAccess) bool {
	value := setting.Data()
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

// SetFollowingsAccess sets Followings access.
func (user *User) SetFollowingsAccess(setting ViewAccess) bool {
	value := setting.Data()
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

// SetFollowersAccess sets Followers access.
func (user *User) SetFollowersAccess(setting ViewAccess) bool {
	value := setting.Data()
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

// SetAllowComment sets AllowComment access.
func (user *User) SetAllowComment(setting ViewAccess) bool {
	value := setting.Data()
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

// SetAllowLike sets AllowLike access.
func (user *User) SetAllowLike(setting ViewAccess) bool {
	value := setting.Data()
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

// SetAllowMessage sets AllowMessage access.
func (user *User) SetAllowMessage(setting ViewAccess) bool {
	value := setting.Data()
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
