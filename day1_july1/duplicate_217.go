package main

func containsDuplicate(nums []int) bool {
	numFreq := make(map[int]int)
	for _, v := range nums {
		numFreq[v]++
	}
	for _, count := range numFreq {
		if count > 1 {
			return true
		}
	}
	return false
}
