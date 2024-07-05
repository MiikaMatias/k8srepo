package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	counter := 0
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, fmt.Sprintf("pong %s", fmt.Sprintf("%v", counter)))
		counter++
	})

	port := os.Getenv("PORT")
	fmt.Printf("Server started in port %s\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatal(err)
	}
}
