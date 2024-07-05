package main

import (
	"fmt"
	"math/rand"
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

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func writeLog(logString string) {
	file, err := os.OpenFile("/logs/log.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	check(err)
	defer file.Close()

	_, err = file.WriteString(logString)
	check(err)

	fileInfo, _ := os.Stat("/logs/log.txt")
	fmt.Printf("Log written successfully: %v\n", fileInfo)
}

func periodicLog(randomString string) {
	for true {
		writeLog(fmt.Sprintf("%s: %s\n", time.Now().Format("2006-01-02 15:04:05"), randomString))
		time.Sleep(5 * time.Second)
	}
}

func main() {
	randomString := randomString()
	go periodicLog(randomString)

	select {}
}
