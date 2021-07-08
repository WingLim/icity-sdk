package data

// Common
const (
	TokenKEY = "authenticity_token"
)

// For login
const (
	Utf8KEY     = "utf8"
	UsernameKEY = "icty_user[login]"
	PasswordKEY = "icty_user[password]"
	CommitKEY   = "commit"
	RememberKEY = "icty_user[remember_me]"

	DefaultUtf8     = "%E2%9C%93"
	DefaultCommit   = "登入"
	DefaultRemember = "1"
)

const (
	MethodKEY = "_method"
)

// For send diary
const (
	TitleKEY   = "activity[data][title]"
	CommentKEY = "activity[data][comment]"
	PrivacyKEY = "activity[data][user_privacy]"
)

const (
	NicknameKey = "user[nickname]"
	BioKey      = "user[info][bio]"
	LocationKey = "user[info][location]"
)
