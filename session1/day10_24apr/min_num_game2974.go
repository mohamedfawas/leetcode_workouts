package main

import (
	"fmt"
	"slices"
)

func numberGame(nums []int) []int {
	arr := make([]int, len(nums))
	slices.Sort(nums)
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			arr[i] = nums[i+1]
		} else {
			arr[i] = nums[i-1]
		}
	}
	return arr
}

func main() {
	nums := []int{5, 4, 2, 3}

	fmt.Println(numberGame(nums))
}
