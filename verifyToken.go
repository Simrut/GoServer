package main

import (
	"github.com/gbrlsnchs/jwt"
	"log"
)

func verifyToken(token []byte) {
	var pl CustomPayload
	hd, err := jwt.Verify(token, hs, &pl)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Print(hd)
		log.Print("Got verified")
	}
}

func main() {

	//TODO implement server that calls verifyToken

}
