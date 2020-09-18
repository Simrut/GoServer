package main

import (
	"github.com/gbrlsnchs/jwt/v3"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type CustomPayload struct {
	jwt.Payload
	Foo string `json:"foo,omitempty"`
	Bar int    `json:"bar,omitempty"`
}

var hs = jwt.NewHS256([]byte("secret"))

// GenerateToken returns a unique token based on the provided email string
func GenerateToken() []byte {
	now := time.Now()
	pl := CustomPayload{
		Payload: jwt.Payload{
			Issuer:         "gbrlsnchs",
			Subject:        "someone",
			Audience:       jwt.Audience{"https://golang.org", "https://jwt.io"},
			ExpirationTime: jwt.NumericDate(now.Add(24 * 30 * 12 * time.Hour)),
			NotBefore:      jwt.NumericDate(now.Add(30 * time.Minute)),
			IssuedAt:       jwt.NumericDate(now),
			JWTID:          "foobar",
		},
		Foo: "foo",
		Bar: 1337,
	}

	token, err := jwt.Sign(pl, hs)
	if err != nil {
		// TODO create error handling
	} else {

		return token
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		//io.WriteString(w, "Hello, TLS!\n")
		io.WriteString(w, GenerateToken(randSeq(10)))
	})

	// One can use generate_cert.go in crypto/tls to generate cert.pem and key.pem.
	log.Printf("About to listen on 8443. Go to http://127.0.0.1:8443/")
	//err := http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", nil)
	err := http.ListenAndServe(":8443", nil)
	log.Fatal(err)
}
