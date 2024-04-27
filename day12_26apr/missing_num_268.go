package main

import (
	"fmt"
)

func missingNumber(nums []int) int {
	var observed_sum int
	for _, v := range nums {
		observed_sum += v
	}
	n := len(nums) // no need to take extra one length, bcz array contains zero value.
	expected_sum := n * (n + 1) / 2
	missing_num := expected_sum - observed_sum
	return missing_num
}

func main() {
	fmt.Println()
}
