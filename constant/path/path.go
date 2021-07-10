package path

const Home = "https://icity.ly"

const (
	Welcome         = "/welcome"
	Login           = "/users/sign_in"
	Logout          = "/users/sign_out"
	World           = "/world"
	Friends         = "/friends"
	UserHome        = "/u/%s"
	UserDiariesPage = "/u/%s/posts?page=%d"

	SetPassword = "/settings/user/password"
)

const (
	Follow   = "/u/%s/follow"
	Unfollow = "/u/%s/unfollow"
)
