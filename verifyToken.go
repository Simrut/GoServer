package main

import (
	"github.com/gbrlsnchs/jwt"
	"log"
)

func verifyToken(token []byte) bool {
	var pl CustomPayload
	hd, err := jwt.Verify(token, hs, &pl)
	if err != nil {
		log.Print("Token incorrect")
		log.Print(err)
		return false
	} else {
		log.Print(hd)
		log.Printf("Got verified")
		return true
	}
}

/*func main() {


}*/
