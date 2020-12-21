package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// KeysURL is the URL containing all the order keys
const KeysURL = "https://www.humblebundle.com/home/keys"

type userOptions struct {
	GameKeys []string `json:"gamekeys"`
}

// GetUserGameKeys gets all the game keys for the user
func GetUserGameKeys(cookie string) ([]string, error) {
	req, err := http.NewRequest("GET", KeysURL, nil)
	if err != nil {
		return nil, err
	}

	req.AddCookie(&http.Cookie{
		Name:     "_simpleauth_sess",
		Value:    cookie,
		Path:     "/",
		Domain:   ".humblebundle.com",
		Secure:   true,
		HttpOnly: true,
	})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	userDataRaw := doc.Find("#user-home-json-data").First().Text()
	if len(userDataRaw) == 0 {
		return nil, fmt.Errorf("Userdata was not found in response, perhaps the cookie is incorrect?")
	}

	var uData userOptions
	err = json.Unmarshal([]byte(userDataRaw), &uData)
	if err != nil {
		return nil, err
	}

	return uData.GameKeys, nil
}
