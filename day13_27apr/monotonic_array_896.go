package main

import "fmt"

func isMonotonic(nums []int) bool {
	isDecreasing, isIncreasing := true, true
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] < nums[i] {
			isIncreasing = false
		} else if nums[i+1] > nums[i] {
			isDecreasing = false
		}
		if !isDecreasing && !isIncreasing {
			return false
		}
	}
	return isDecreasing || isIncreasing
}

func main() {
	fmt.Println()
}
