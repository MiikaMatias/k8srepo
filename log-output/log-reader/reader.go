package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

const (
	hashFilePath    = "/logs/log.txt"
	configFilePath  = "/config/log.properties"
	messageFilePath = "/config/message"
	pingPongUrl     = "ping-pong-svc:2345/no-increment"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hash, err := os.ReadFile(hashFilePath)
		check(err)
		configText, err := os.ReadFile(configFilePath)
		messageText, err := os.ReadFile(messageFilePath)
		cmd := exec.Command("curl", pingPongUrl)
		pingpong, err := cmd.Output()
		check(err)
		fmt.Fprintf(w, fmt.Sprintf("file content: %s\nenv variable: %s \n%s \nPings / Pongs: %s",
			string(configText), string(messageText),
			string(hash), string(pingpong[len(pingpong)-1])))
	})

	port := os.Getenv("PORT")
	fmt.Printf("Server started in port %s\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatal(err)
	}
}
