package path

const HOME = "https://icity.ly"

const (
	WELCOME = "/welcome"
	SIGNIN  = "/users/sign_in"
	SIGNOUT = "/users/sign_out"
	WORLD   = "/world"
)

// Diary Url Path
const (
	NEWDIARY    = "/activities?act_id=301&re_id=icty%2Buser%2B1irqw29"
	DELETEDIARY = "/a/%s?act_id=301&re_id=icty%%2Buser%%2B1irqw29"

	LIKE   = "/a/%s/like"
	UNLIKE = "/a/%s/unlike"

	SETPRIVACY = "/a/%s/set_privacy?user_privacy=%d"

	COMMENT = "/activities?act_id=101&re_id=icty%%2Bactivity%%2B%s"
)
