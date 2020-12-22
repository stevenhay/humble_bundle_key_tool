package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print("_simpleauth_sess -> ")
	var cookieInput string
	fmt.Scanln(&cookieInput)

	// change \075 for = if it exists
	cookieInput = strings.Replace(cookieInput, "\\075", "=", 1)

	client := NewClient(cookieInput)

	humblebundle := NewHumbleBundle(client)
	gamekeys, err := humblebundle.GetUserGameKeys()
	if err != nil {
		fmt.Printf("[FAILED] %s", err.Error())
		return
	}

	infoChan := make(chan *GameKeyInformation)
	for _, s := range gamekeys {
		go func(s string) {
			info, err := humblebundle.GetInformationForGameKey(s)
			if err != nil {
				fmt.Printf("[FAILED] %s", err.Error())
				return
			}
			infoChan <- info
		}(s)
	}

	expiredKeys := make([]Tpkd, 0)
	unredeemedKeys := make([]Tpkd, 0)
	for i := 0; i < len(gamekeys); i++ {
		v := <-infoChan
		for _, game := range v.TpkdDict.AllTpks {
			if game.IsExpired {
				expiredKeys = append(expiredKeys, game)
			} else if len(game.RedeemedKeyValue) == 0 {
				unredeemedKeys = append(unredeemedKeys, game)
			}
		}
	}

	fmt.Println("--- EXPIRED GAMES ---")
	for _, v := range expiredKeys {
		fmt.Printf("[%s] %s\n", v.KeyType, v.HumanName)
	}

	fmt.Println("--- UNREDEEMED GAMES ---")
	for _, v := range unredeemedKeys {
		fmt.Printf("[%s] %s\n", v.KeyType, v.HumanName)
	}
}
