package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	// Create the routing
	r := routing()

	// Bind the router to the http module
	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("server is running")
	log.Fatal(srv.ListenAndServe())
}
