package icity_sdk

var (
	defaultOption = &Options{
		SaveCookies: false,
	}
)

type Option func(*Options)

type Options struct {
	SaveCookies bool
}

func WithSaveCookies() Option {
	return func(o *Options) {
		o.SaveCookies = true
	}
}
