package main

import "sort"

func minOperations(nums []int, k int) int {
	sort.Ints(nums)
	counter := 0
	for _, num := range nums {
		if num < k {
			if counter <= k {
				counter++
			} else {
				break
			}
		}
	}

	return counter
}
