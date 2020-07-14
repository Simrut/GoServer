package main

import (
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/sha3"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890./!#_")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// GenerateToken returns a unique token based on the provided email string
func GenerateToken(randomstring string) string {
	hasher := sha3.New512()
	hash, err := hasher.Write([]byte(randomstring))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hash to store:", string(hash))

	return hex.EncodeToString(hasher.Sum(nil))
}

func main() { //TODO switch to JWT?
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		randomseq := randSeq(20)
		fmt.Println(randomseq)
		io.WriteString(w, GenerateToken(randomseq))
	})

	// One can use generate_cert.go in crypto/tls to generate cert.pem and key.pem.
	log.Printf("About to listen on 8443. Go to http://127.0.0.1:8443/")
	//err := http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", nil)
	err := http.ListenAndServe(":8443", nil)
	log.Fatal(err)
}
