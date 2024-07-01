package main

import (
	"fmt"
	"slices"
)

func arrayPairSum(nums []int) int {
	slices.Sort(nums)
	var sum int
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			sum += min(nums[i], nums[i+1])
		}
	}
	return sum
}

func main() {
	arr := []int{1, 4, 3, 2}
	fmt.Println(arrayPairSum(arr))
}
