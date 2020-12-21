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
	gamekeys, err := GetUserGameKeys(client)
	if err != nil {
		fmt.Printf("[FAILED] %s", err.Error())
		return
	}

	for _, s := range gamekeys {
		fmt.Println(s)
	}
}
