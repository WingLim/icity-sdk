package icity_sdk

import (
	"log"
	"net/http"
	"net/http/cookiejar"
)

type User struct {
	Username string
	Password string

	token  string
	client http.Client
}

func NewUser(username, password string) *User {
	user := &User{
		Username: username,
		Password: password,
	}
	user.initUser()
	return user
}

func (user *User) initUser() {
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal(err)
	}
	user.client.Jar = jar
}
