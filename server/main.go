package main

import (
	"fmt"
	"log"
	"net/http"

	server "github.com/owen-corrigan-bjss/to-do-app/server/server"
)

func main() {
	server.NewServer()
	fmt.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
