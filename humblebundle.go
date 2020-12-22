package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const keysURL = "https://www.humblebundle.com/home/keys"
const gameKeyURL = "https://www.humblebundle.com/api/v1/order/%s?all_tpkds=true"

type userOptions struct {
	GameKeys []string `json:"gamekeys"`
}

// GameKeyInformation contains all information about a purchase
type GameKeyInformation struct {
	Product          GameKeyProduct `json:"product"`
	TpkdDict         TpkdDictionary `json:"tpkd_dict"`
	TotalChoices     int            `json:"total_choices"`
	RemainingChoices int            `json:"choices_remaining"`
}

// GameKeyProduct contains information about the type of thing purchased,
// for example, which bundle it was.
type GameKeyProduct struct {
	IsHumbleChoice bool   `json:"is_humble_choice"`
	MachineName    string `json:"machine_name"`
	HumanName      string `json:"human_name"`
}

// TpkdDictionary contains a list of Tpkd, not sure what this stands for yet.
type TpkdDictionary struct {
	AllTpks []Tpkd `json:"all_tpks"`
}

// Tpkd contains information about games within a bundle, such as the steam key (if it's
// been redeemed), if it's expired, the name of the game, etc. It does *not* include information
// on unredeemed games that are part of the monthly bundles.
type Tpkd struct {
	MachineName      string `json:"machine_name"`
	HumanName        string `json:"human_name"`
	KeyType          string `json:"key_type"`
	RedeemedKeyValue string `json:"redeemed_key_val"` // any tpkd missing this value is unredeemed
	IsExpired        bool   `json:"is_expired"`
}

// HumbleBundle is used for interacting with the HumbleBundle website
type HumbleBundle struct {
	client Client
}

// NewHumbleBundle creates a new struct with a client containing a session cookie
func NewHumbleBundle(client Client) HumbleBundle {
	return HumbleBundle{client}
}

// GetUserGameKeys gets all the game keys for the user
func (h HumbleBundle) GetUserGameKeys() ([]string, error) {
	resp, err := h.client.Do("GET", keysURL, nil)
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

// GetInformationForGameKey returns information for a gamekey, this information
// may contain many games for a single key. The 'gamekey' is more like a purchase
// key as purchases can contain more than one game in a bundle.
func (h HumbleBundle) GetInformationForGameKey(gamekey string) (*GameKeyInformation, error) {
	resp, err := h.client.Do("GET", fmt.Sprintf(gameKeyURL, gamekey), nil)
	if err != nil {
		return nil, err
	}

	var gameKeyInfo GameKeyInformation
	err = json.Unmarshal(resp, &gameKeyInfo)
	if err != nil {
		return nil, err
	}

	return &gameKeyInfo, nil
}
