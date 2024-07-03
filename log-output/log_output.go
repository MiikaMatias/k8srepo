package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

const (
	letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789-"
	hashlen = 20
)

func randomString() string {
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, hashlen)

	for i := range hashlen {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func periodicLog(currentState *string, randomString string) {
	for true {
		*currentState = fmt.Sprintf("%s: %s\n", time.Now().Format("2006-01-02 15:04:05"), randomString)
		time.Sleep(5 * time.Second)
	}
}

func main() {
	randomString := randomString()
	currentState := ""
	go periodicLog(&currentState, randomString)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, currentState)
	})

	port := os.Getenv("PORT")
	fmt.Printf("Server started in port %s\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatal(err)
	}
}
