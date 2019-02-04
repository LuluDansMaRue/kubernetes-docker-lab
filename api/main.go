package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Create the routing
	r := routing()

	fmt.Println("server is running")

	// Bind the router to the http module
	log.Fatal(http.ListenAndServe(":8000", r))
}
