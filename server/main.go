package main

import (
	"fmt"
	"log"
	"net/http"
	handlers "to-do-app/server/handlers"
)

func main() {
	handlers.Handlers()
	fmt.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
