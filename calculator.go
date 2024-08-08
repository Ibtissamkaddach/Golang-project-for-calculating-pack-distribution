package calculator

// CalculatePacks calculates the number of packs needed for a given quantity
func CalculatePacks(items int) map[int]int {
	packSizes := []int{5000, 2000, 1000, 500, 250}
	packs := make(map[int]int)
	remainingItems := items

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

	if remainingItems > 0 {
		smallestPack := packSizes[len(packSizes)-1]
		packs[smallestPack]++
	}

	return packs
}
