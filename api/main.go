package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/rs/cors"
)

func main() {
	// Create the routing
	r := routing()

	// Create a handler for the CORS
	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE"},
	}).Handler(r)

	// Bind the router to the http module
	srv := &http.Server{
		Handler:      handler,
		Addr:         "0.0.0.0:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("server is running")
	log.Fatal(srv.ListenAndServe())
}
