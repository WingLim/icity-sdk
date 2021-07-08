package selector

const (
	LoginToken = "input[name=authenticity_token]"
	CSRFToken  = "meta[name=csrf-token]"
	WORLD      = "body > div.container.below-top-navbar > div > div.box-head > h2"

	NICKNAME = "#user_nickname"
	BIO      = "#user_info_bio"
	LOCATION = "#user_info_location"
)

const (
	WorldDiarys   = "ul.posts-list li.wu"
	DiaryId       = "a.timeago"
	DiaryNickname = "a.user"
	DiaryTitle    = "div.line > h4 > a"
	DiaryContent  = "div.comment"
	DiaryLocation = "span.location"
	DiaryDate     = "time.timeago"

	Comments    = "div.cntr > div > div.cntr > ul > li.wu"
	CommentDate = "time.hours"
)

const (
	optionSelected         = " option[selected=selected]"
	SettingsDefaultPrivacy = "#user_settings_default_privacy" + optionSelected
	SettingsPrivacy        = "#user_settings_privacy" + optionSelected
	SettingsUnWorld        = "#user_settings_user_unworld[checked=checked]"
	SettingsAboutMe        = "#user_settings_privacy_about_me" + optionSelected
	SettingsFollowings     = "#user_settings_privacy_followings" + optionSelected
	SettingsFollowers      = "#user_settings_privacy_followers" + optionSelected
	SettingsAllowComment   = "#user_settings_privacy_allow_comment" + optionSelected
	SettingsAllowLike      = "#user_settings_privacy_allow_like" + optionSelected
	SettingsAllowMessage   = "#user_settings_privacy_allow_message" + optionSelected
)
