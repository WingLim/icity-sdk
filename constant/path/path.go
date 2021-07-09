package path

const Home = "https://icity.ly"

const (
	Welcome = "/welcome"
	Login   = "/users/sign_in"
	Logout  = "/users/sign_out"
	World   = "/world"
)

// Diary Url Path
const (
	NewDiary    = "/activities?act_id=301&re_id=icty%2Buser%2B1irqw29"
	DeleteDiary = "/a/%s?act_id=301&re_id=icty%%2Buser%%2B1irqw29"

	Like   = "/a/%s/like"
	Unlike = "/a/%s/unlike"

	DiaryPrivacy = "/a/%s/set_privacy?user_privacy=%d"

	NewComment    = "/activities?act_id=101&re_id=icty%%2Bactivity%%2B%s"
	DeleteComment = "/a/%s?act_id=101&re_id=icty%%2Bactivity%%2B%s"
	GetComments   = "/a/%s/expand"
)

const (
	Follow   = "/u/%s/follow"
	Unfollow = "/u/%s/unfollow"
)

const (
	SettingsIndex = "/settings/user/index"
	Social        = "/settings/user/social"
	Privacy       = "/settings/user/privacy"
)

const (
	Conversations = "/user/conversations"
	Conversation  = "/user/conversations/%s"
)
