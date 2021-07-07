package icity_sdk

import (
	"net/http"
)

type User struct {
	Username string
	Password string

	client http.Client
}

func NewUser(username, password string) *User {
	user := &User{
		Username: username,
		Password: password,
	}
	user.initClient()
	return user
}
