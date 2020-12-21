package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// KeysURL is the URL containing all the order keys
const KeysURL = "https://www.humblebundle.com/home/keys"

type userOptions struct {
	GameKeys []string `json:"gamekeys"`
}

// GetUserGameKeys gets all the game keys for the user
func GetUserGameKeys(client Client) ([]string, error) {
	resp, err := client.Do("GET", KeysURL, nil)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(resp)))
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
