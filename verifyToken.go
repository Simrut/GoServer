package main

import (
	"github.com/gbrlsnchs/jwt/v3"
	"log"
)

func ReceiveToken() []byte {
	return []byte("Test")
}

func main() {

	token := ReceiveToken()
	var pl CustomPayload
	hd, err := jwt.Verify(token, hs, &pl)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Print(hd)
		log.Print("Got verified")
	}

}
