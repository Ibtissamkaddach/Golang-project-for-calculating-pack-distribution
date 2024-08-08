package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// packHandler handles the request to calculate packs
func packHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		quantityStr := r.FormValue("quantity")
		quantity, err := strconv.Atoi(quantityStr)
		if err != nil {
			http.Error(w, "Invalid quantity", http.StatusBadRequest)
			return
		}

		// Calculate packs
		packs := calculatePacks(quantity)

		// Write the result as a response
		fmt.Fprintf(w, "Packs: %v\n", packs)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
