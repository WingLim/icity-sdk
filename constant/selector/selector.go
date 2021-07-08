package selector

const (
	LoginToken = "input[name=authenticity_token]"
	CSRFToken  = "meta[name=csrf-token]"
	WORLD      = "body > div.container.below-top-navbar > div > div.box-head > h2"
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
