package main

import (
	"flag"
	"log"

	"github.com/sinjs/clicord/cmd"
	"github.com/sinjs/clicord/internal/constants"
	"github.com/zalando/go-keyring"
)

func main() {
	// Declare and parse all flags first
	token := flag.String("token", "", "authentication token")
	flag.Parse()

	// If no token was provided, look it up in the keyring
	if *token == "" {
		t, err := keyring.Get(constants.Name, "token")
		if err != nil {
			log.Println("failed to get token from keyring:", err)
		} else {
			*token = t
		}
	}

	if err := cmd.Run(*token); err != nil {
		log.Fatal(err)
	}
}
