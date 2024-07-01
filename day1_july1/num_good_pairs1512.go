package main

func numIdenticalPairs(nums []int) int {
	var pairCount int
	numFreq := make(map[int]int)
	for _, v := range nums {
		numFreq[v]++
	}
	for _, count := range numFreq {
		pairCount += count * (count - 1) / 2
	}
	return pairCount
}
