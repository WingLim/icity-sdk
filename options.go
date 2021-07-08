package icity_sdk

import "net/http"

var (
	defaultOption = &Options{
		Cookies:     nil,
		CookiesPath: "",
	}
)

type Option func(*Options)

type Options struct {
	Cookies     []*http.Cookie
	CookiesPath string
}

func WithSaveCookies(filepath string) Option {
	return func(o *Options) {
		o.Cookies = readCookiesFromFile(filepath)
		o.CookiesPath = filepath
	}
}
