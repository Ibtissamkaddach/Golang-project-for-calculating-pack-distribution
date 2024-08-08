package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/calculate", packHandler) // Register the handler for /calculate
	r := server.NewRouter()
	port := "8080"
	log.Printf("Starting server on port %s", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
