package main

import (
	"log"
	"net/http"
)

func main() {
	// Set up routes
	http.HandleFunc("/calculate", packHandler)
	// Start the server
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
