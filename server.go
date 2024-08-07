package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// packHandler handles HTTP requests to calculate pack distribution
func packHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the number of items from query parameters
	r.ParseForm()
	itemsStr := r.FormValue("items")
	items, err := strconv.Atoi(itemsStr)
	if err != nil || items <= 0 {
		http.Error(w, "Invalid number of items", http.StatusBadRequest)
		return
	}

	// Calculate the number of packs needed
	packs, err := calculatePacks(items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Return the result as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(packs)
}
