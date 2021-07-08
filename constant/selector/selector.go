package selector

const (
	LoginToken = "input[name=authenticity_token]"
	CSRFToken  = "meta[name=csrf-token]"
	WORLD      = "body > div.container.below-top-navbar > div > div.box-head > h2"
)

const (
	WorldDiarys   = "ul.posts-list li.wu"
	DiaryId       = "div.meta > a.timeago"
	DiaryNickname = "div.line > a"
	DiaryTitle    = "div.line > h4 > a"
	DiaryContent  = "div.line > div.comment"
	DiaryLocation = "div.line > span.location"
	DiaryDate     = "div.meta > a.timeago > time"

	Comments    = "div.cntr > div > div.cntr > ul > li.wu"
	CommentDate = "time.hours"
)
