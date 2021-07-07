package icity_sdk

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func closeBody(body io.ReadCloser) {
	if err := body.Close(); err != nil {
		log.Fatal(err)
	}
}

func saveCookiesToFile(cookies []*http.Cookie) error {
	f, err := os.Create("cookies.json")
	if err != nil {
		return err
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	err = encoder.Encode(cookies)
	if err != nil {
		return err
	}
	return nil
}

func readCookiesFromFile() []*http.Cookie {
	f, err := os.Open("cookies.json")
	if err != nil {
		return nil
	}
	defer f.Close()

	var cookies []*http.Cookie
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&cookies)
	if err != nil {
		return nil
	}
	return cookies
}
