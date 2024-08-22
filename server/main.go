package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type ToDo struct {
	description string
	completed   bool
	lock        sync.Mutex
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleGetHello(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleGetHello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("hi")
}

func main() {
	fmt.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
