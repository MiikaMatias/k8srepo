package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

type Todo struct {
	Todos []string `json:"todos"`
}

var todoList []string
var mutex sync.Mutex

func main() {
	http.HandleFunc("/", handleRequest)
	port := os.Getenv("PORT")
	if port == "" {
		log.Print("port not set, setting to 80")
		port = "80"
	}
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatal(err)
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	if r.Method == "POST" {
		var todos Todo
		err := json.NewDecoder(r.Body).Decode(&todos)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		todoList = append(todoList, todos.Todos...)
	}
	if r.Method == "GET" {
		wrappedData := Todo{
			Todos: todoList,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode((wrappedData))
	}
}
