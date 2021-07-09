package selector

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

	Devices          = "table tbody tr"
	DeviceApp        = "td:nth-child(1)"
	DeviceHardware   = "td:nth-child(1) > small"
	DeviceLastUsed   = "td:nth-child(2) > small"
	DeviceRemoveLink = "td:nth-child(3) > a"
)
