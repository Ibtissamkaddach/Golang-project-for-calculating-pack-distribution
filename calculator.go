package main

// calculatePacks determines the number of packs needed for the given number of items.
func calculatePacks(items int) map[int]int {
	// Define available pack sizes in descending order
	packSizes := []int{5000, 2000, 1000, 500, 250}
	packs := make(map[int]int)
	remainingItems := items

	// Iterate through pack sizes and calculate the number of each pack needed
	for _, size := range packSizes {
		if remainingItems == 0 {
			break
		}
		count := remainingItems / size
		if count > 0 {
			packs[size] = count
			remainingItems -= size * count
		}
	}

	// If there are remaining items, add one more pack of the smallest size
	if remainingItems > 0 {
		smallestPack := packSizes[len(packSizes)-1]
		packs[smallestPack]++
	}

	return packs
}
