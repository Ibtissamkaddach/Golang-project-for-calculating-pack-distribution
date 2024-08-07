package main

import (
    "encoding/json"
    "log"
    "net/http"
    "sort"
    "strconv"
)

// Config structure to hold pack sizes
type Config struct {
    PackSizes []int `json:"pack_sizes"`
}

// Global configuration
var config = Config{
    PackSizes: []int{250, 500, 1000, 2000, 5000},
}

// Function to calculate packs
func calculatePacks(order int) map[int]int {
    result := make(map[int]int)
    remaining := order

    // Sort pack sizes in descending order
    sort.Sort(sort.Reverse(sort.IntSlice(config.PackSizes)))

    for _, pack := range config.PackSizes {
        if remaining == 0 {
            break
        }
        if remaining >= pack {
            numPacks := remaining / pack
            result[pack] = numPacks
            remaining -= pack * numPacks
        }
    }

    // If there are any remaining items, add the smallest pack
    if remaining > 0 {
        result[config.PackSizes[len(config.PackSizes)-1]]++
    }

    return result
}

// HTTP handler for pack calculation
func packHandler(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query()
    orderStr := query.Get("order")
    if orderStr == "" {
        http.Error(w, "Missing order parameter", http.StatusBadRequest)
        return
    }

    order, err := strconv.Atoi(orderStr)
    if err != nil {
        http.Error(w, "Invalid order parameter", http.StatusBadRequest)
        return
    }

    packs := calculatePacks(order)
    response, err := json.Marshal(packs)
    if err != nil {
        http.Error(w, "Error processing request", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(response)
}

func main() {
    http.HandleFunc("/calculate-packs", packHandler)
    log.Println("Server is running on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}