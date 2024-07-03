package main

import (
	"fmt"
	"math/rand"
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

func main() {
	randomString := randomString()
	for true {
		fmt.Printf("%s: %s\n", time.Now().Format("2006-01-02 15:04:05"), randomString)
		time.Sleep(5 * time.Second)
	}
}
