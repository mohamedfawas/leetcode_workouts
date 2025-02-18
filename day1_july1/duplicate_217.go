package main

func containsDuplicate(nums []int) bool {
	numRecord := make(map[int]bool)
	for _, num := range nums {
		if numRecord[num] {
			return true
		} else {
			numRecord[num] = true
		}
	}

	return false
}
