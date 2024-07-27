package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

const (
	logFilePath = "/logs/log.txt"
	pingPongUrl = "ping-pong-svc:2345/no-increment"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logs, err := os.ReadFile(logFilePath)
		check(err)
		cmd := exec.Command("curl", pingPongUrl)
		pingpong, err := cmd.Output()
		check(err)
		fmt.Fprintf(w, fmt.Sprintf("%s\nPings / Pongs: %s", string(logs), string(pingpong[len(pingpong)-1])))
	})

	port := os.Getenv("PORT")
	fmt.Printf("Server started in port %s\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatal(err)
	}
}
