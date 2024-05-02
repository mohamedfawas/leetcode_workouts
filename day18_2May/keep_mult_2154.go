package main

import (
	"fmt"
	"slices"
)

func findFinalValue(nums []int, original int) int {
	slices.Sort(nums)
	for i := 0; i < len(nums); i++ {
		if original == nums[i] {
			original *= 2
		}
	}
	return original
}

func main() {
	fmt.Println
}
