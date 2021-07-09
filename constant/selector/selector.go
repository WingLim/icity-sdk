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
