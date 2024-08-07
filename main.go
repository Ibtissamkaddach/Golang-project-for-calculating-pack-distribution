package main

import (
	"log"
	"net/http"
)

func main() {
	// Initialize the server and define routes
	http.HandleFunc("/calculate", packHandler)
	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
