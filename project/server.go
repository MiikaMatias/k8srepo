package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	mainImageUrl  = "https://picsum.photos/1200"
	mainImagePath = "assets/main-image.jpg"
)

func fetchImage() {
	for {
		response, err := http.Get(mainImageUrl)
		if err != nil {
			log.Printf("Failed to fetch image: %v", err)
			continue
		}
		if response.StatusCode != http.StatusOK {
			log.Printf("Non-OK HTTP status: %d", response.StatusCode)
			response.Body.Close()
			continue
		}
		file, err := os.Create(mainImagePath)
		if err != nil {
			log.Printf("Failed to create file: %v", err)
			response.Body.Close()
			continue
		}
		_, err = io.Copy(file, response.Body)
		if err != nil {
			log.Printf("Failed to save image: %v", err)
			continue
		}
		log.Println("Fetched a new image")
		file.Close()
		response.Body.Close()
		time.Sleep(60 * time.Minute)
	}
}

func serveHtml(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func main() {
	go fetchImage()

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/", serveHtml)

	port := os.Getenv("PORT")
	if port == "" {
		log.Print("port not set, setting to 80")
		port = "80"
	}
	fmt.Printf("Server started in port %s\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatal(err)
	}
}
