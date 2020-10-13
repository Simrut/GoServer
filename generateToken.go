package main

import (
	"github.com/gbrlsnchs/jwt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

type CustomPayload struct {
	jwt.Payload
	RandSeq string `json:"foo,omitempty"`
}

var hs = jwt.NewHS256([]byte("lgg3d5sf8v3"))

func GenerateToken() []byte {
	now := time.Now()
	pl := CustomPayload{
		Payload: jwt.Payload{
			Issuer:         "theinsect",
			Subject:        "transmission",
			Audience:       jwt.Audience{"https://golang.org", "https://jwt.io"},
			ExpirationTime: jwt.NumericDate(now.Add(24 * 30 * 12 * time.Hour)),
			NotBefore:      jwt.NumericDate(now.Add(30 * time.Minute)),
			IssuedAt:       jwt.NumericDate(now),
			JWTID:          "f5rek432",
		},
		RandSeq: randSeq(20),
	}

	token, err := jwt.Sign(pl, hs)
	if err != nil {
		log.Fatal(err)
	}
	return token
}

func main() {
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		token := GenerateToken()
		io.WriteString(w, string(token))
	})

	// One can use generate_cert.go in crypto/tls to generate cert.pem and key.pem.
	log.Printf("About to listen on 8443. Go to https://127.0.0.1:8443/")
	err := http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", nil)

	//err := http.ListenAndServe(":8443", nil)
	log.Fatal(err)

}
