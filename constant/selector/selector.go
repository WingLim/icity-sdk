package selector

const (
	LoginToken = "input[name=authenticity_token]"
	CSRFToken  = "meta[name=csrf-token]"
	World      = "body > div.container.below-top-navbar > div > div.box-head > h2"

	UserID   = "div.navbar.navbar-major.fixedsticky > div > ul > li:nth-child(4) > a"
	Nickname = "#user_nickname"
	Bio      = "#user_info_bio"
	Location = "#user_info_location"
)

const (
	WorldDiaries  = "ul.posts-list li.wu"
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

const (
	ConversationsList           = "ul.conversations-list li"
	ConversationItemID          = "li > a"
	ConversationItemUser        = "span.user"
	ConversationItemLastMessage = "div.line"
	ConversationItemLastDate    = "time.timeago"

	Conversations       = "ul.messages-list li"
	ConversationContent = "div.box"
)
