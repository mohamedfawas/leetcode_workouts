package main

import "math"

func differenceOfSum(nums []int) int {
	var el_sum, dig_sum, temp_val int
	for i := 0; i < len(nums); i++ {
		el_sum += nums[i]
		temp_val = nums[i]
		for temp_val > 0 {
			l_digit := temp_val % 10
			dig_sum += l_digit
			temp_val /= 10
		}
	}
	return int(math.Abs(float64(el_sum) - float64(dig_sum)))
}
