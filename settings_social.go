package icity_sdk

import (
	"net/http"

	"github.com/WingLim/icity-sdk/constant/data"
	"github.com/WingLim/icity-sdk/constant/path"
	"github.com/WingLim/icity-sdk/log"
)

func (user *User) doSocialSettings(settings ...Setting) bool {
	postData := buildSettingsData(settings)

	resp, err := user.postForm(path.SOCIAL, postData)
	if err != nil {
		log.Error(err)
		return false
	}

	if resp.StatusCode == http.StatusOK {
		return true
	}
	return false
}

func (user *User) SetSocialWeibo(weibo string) bool {
	set := Setting{
		Key:   data.WeiboKey,
		Value: weibo,
	}
	return user.doSocialSettings(set)
}

func (user *User) SetSocialWechat(wechat string) bool {
	set := Setting{
		Key:   data.WechatKey,
		Value: wechat,
	}
	return user.doSocialSettings(set)
}

func (user *User) SetSocialQQ(qq string) bool {
	set := Setting{
		Key:   data.QQKey,
		Value: qq,
	}
	return user.doSocialSettings(set)
}

func (user *User) SetSocialInstagram(instagram string) bool {
	set := Setting{
		Key:   data.InstagramKey,
		Value: instagram,
	}
	return user.doSocialSettings(set)
}

func (user *User) SetSocialTwitter(twitter string) bool {
	set := Setting{
		Key:   data.TwitterKey,
		Value: twitter,
	}
	return user.doSocialSettings(set)
}

func (user *User) SetSocialFacebook(facebook string) bool {
	set := Setting{
		Key:   data.FacebookKey,
		Value: facebook,
	}
	return user.doSocialSettings(set)
}

func (user *User) SetSocialTelegram(telegram string) bool {
	set := Setting{
		Key:   data.TelegramKey,
		Value: telegram,
	}
	return user.doSocialSettings(set)
}

func (user *User) SetSocialTumblr(tumblr string) bool {
	set := Setting{
		Key:   data.TumblrKey,
		Value: tumblr,
	}
	return user.doSocialSettings(set)
}

func (user *User) SetSocialGithub(github string) bool {
	set := Setting{
		Key:   data.GithubKey,
		Value: github,
	}
	return user.doSocialSettings(set)
}

func (user *User) SetSocialDribbble(dribbble string) bool {
	set := Setting{
		Key:   data.DribbbleKey,
		Value: dribbble,
	}
	return user.doSocialSettings(set)
}

func (user *User) SetSocialPixiv(pixiv string) bool {
	set := Setting{
		Key:   data.PixivKey,
		Value: pixiv,
	}
	return user.doSocialSettings(set)
}

func (user *User) SetSocialSteam(steam string) bool {
	set := Setting{
		Key:   data.SteamKey,
		Value: steam,
	}
	return user.doSocialSettings(set)
}

func (user *User) SetSocialBattlenet(battlenet string) bool {
	set := Setting{
		Key:   data.BattlenetKey,
		Value: battlenet,
	}
	return user.doSocialSettings(set)
}

func (user *User) SetSocialPSN(psn string) bool {
	set := Setting{
		Key:   data.PSNKey,
		Value: psn,
	}
	return user.doSocialSettings(set)
}

func (user *User) SetSocialNintendo(nintendo string) bool {
	set := Setting{
		Key:   data.NintendoKey,
		Value: nintendo,
	}
	return user.doSocialSettings(set)
}

func (user *User) SetSocialNS(ns string) bool {
	set := Setting{
		Key:   data.NSKey,
		Value: ns,
	}
	return user.doSocialSettings(set)
}

func (user *User) SetSocialXBox(xbox string) bool {
	set := Setting{
		Key:   data.XBoxKey,
		Value: xbox,
	}
	return user.doSocialSettings(set)
}

func (user *User) SetSocialHomePage(homepage string) bool {
	set := Setting{
		Key:   data.HomePageKey,
		Value: homepage,
	}
	return user.doSocialSettings(set)
}

func (user *User) SetSocialLink1(link1 string) bool {
	set := Setting{
		Key:   data.Link1Key,
		Value: link1,
	}
	return user.doSocialSettings(set)
}

func (user *User) SetSocialLink2(link2 string) bool {
	set := Setting{
		Key:   data.Link2Key,
		Value: link2,
	}
	return user.doSocialSettings(set)
}

func (user *User) SetSocialLink3(link3 string) bool {
	set := Setting{
		Key:   data.Link3Key,
		Value: link3,
	}
	return user.doSocialSettings(set)
}
