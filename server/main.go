package main

import (
	"fmt"
	handlers "github.com/owen-corrigan-bjss/to-do-app/server/handlers"
	"log"
	"net/http"
)

func main() {
	handlers.Handlers()
	fmt.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
