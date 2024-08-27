package main

import (
	"fmt"
	"log"
	"net/http"

	handlers "github.com/owen-corrigan-bjss/to-do-app/server/handlers"
)

func main() {
	handlers.NewServer()
	fmt.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
