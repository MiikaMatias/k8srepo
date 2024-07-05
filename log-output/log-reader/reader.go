package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	logFilePath     = "/logs/log.txt"
	pingLogFilePath = "/pinglogs/pinglog.txt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logs, err := os.ReadFile(logFilePath)
		check(err)
		pingpong, err := os.ReadFile(pingLogFilePath)
		check(err)
		fmt.Fprintf(w, fmt.Sprintf("%s\nPings / Pongs: %s", string(logs), string(pingpong)))
	})

	port := os.Getenv("PORT")
	fmt.Printf("Server started in port %s\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatal(err)
	}
}
