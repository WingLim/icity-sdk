package icity_sdk

type InvisibleType int

const (
	NotInvisible InvisibleType = iota + 1
	ToStrangers
	ToAll
)

type ViewAccess int

const (
	Everyone ViewAccess = iota + 1
	Friend
	Self
)

type WorldType int

const (
	Show WorldType = iota
	NoShow
)

type SettingsPrivacy struct {
	DefaultPrivacy DiaryPrivacy
	InvisibleMode  InvisibleType
	UnWorld        WorldType

	AboutMe        ViewAccess
	FollowingsList ViewAccess
	FollowersList  ViewAccess

	AllowComment ViewAccess
	AllowLike    ViewAccess
	AllowMessage ViewAccess
}
