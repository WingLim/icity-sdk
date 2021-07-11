package icity_sdk

import (
	"net/http"
	"net/url"

	"github.com/WingLim/icity-sdk/constant/data"
	"github.com/WingLim/icity-sdk/constant/path"
	"github.com/WingLim/icity-sdk/log"
)

type Setting struct {
	Key   string
	Value string
}

func buildSettingsData(settings []Setting) url.Values {
	postData := url.Values{}
	postData.Set(data.MethodKEY, "put")

	for _, one := range settings {
		postData.Set(one.Key, one.Value)
	}
	return postData
}

func (user *User) doInfoSettings(settings ...Setting) bool {
	postData := buildSettingsData(settings)

	resp, err := user.postForm(path.SettingsIndex, postData)
	if err != nil {
		log.Error(err)
		return false
	}

	if resp.StatusCode == http.StatusOK {
		return true
	}
	return false
}

// SetMyNickName sets user nickname.
func (user *User) SetMyNickName(nickname string) bool {
	set := Setting{
		Key:   data.NicknameKey,
		Value: nickname,
	}
	return user.doInfoSettings(set)
}

// SetMyBio sets user bio.
func (user *User) SetMyBio(bio string) bool {
	set := Setting{
		Key:   data.BioKey,
		Value: bio,
	}
	return user.doInfoSettings(set)
}

// SetMyLocation sets user location.
func (user *User) SetMyLocation(location string) bool {
	set := Setting{
		Key:   data.LocationKey,
		Value: location,
	}
	return user.doInfoSettings(set)
}
