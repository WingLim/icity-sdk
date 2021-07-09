package path

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
