package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	pingLogFilePath = "/pinglogs/pinglog.txt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func writePing(pingString string) {
	fmt.Printf("Started open file")
	file, err := os.OpenFile(pingLogFilePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	check(err)
	defer file.Close()

	fmt.Printf("Started write ping")
	_, err = file.WriteString(pingString)
	check(err)

	fileInfo, _ := os.Stat(pingLogFilePath)
	fmt.Printf("Log written successfully: %v\n", fileInfo)
}

func main() {
	// take this retarded variable and just use the file in the volume
	counter := 0
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, fmt.Sprintf("pong %s", fmt.Sprintf("%v", counter)))
		writePing(fmt.Sprintf("%v", counter))
		counter++
	})
	http.HandleFunc("/no-increment", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, fmt.Sprintf("pong %s", fmt.Sprintf("%v", counter)))
	})

	port := os.Getenv("PORT")
	fmt.Printf("Server started in port %s\n", port)
	writePing("0")
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatal(err)
	}
}
