package main

func countPairs(nums []int, target int) int {
	var count int
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i]+nums[j] < target {
				count++
			}
		}
	}
	return count
}
