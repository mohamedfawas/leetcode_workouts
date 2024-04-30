package main

import (
	"fmt"
	"math"
)

func differenceOfSum(nums []int) int {
	var elem_sum, digit_sum int
	for i := 0; i < len(nums); i++ {
		elem_sum += nums[i]
		temp_val := nums[i]
		for temp_val > 0 {
			digit_sum += temp_val % 10
			temp_val /= 10
		}
	}
	abs_diff := int(math.Abs(float64(digit_sum) - float64(elem_sum)))
	return abs_diff
}

func main() {
	fmt.Println()
}
