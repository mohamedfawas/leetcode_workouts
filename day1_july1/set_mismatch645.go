package main

import "fmt"

func findErrorNums(nums []int) []int {
	n := len(nums)
	seen := make([]bool, n+1)
	var sum, duplicate, missing int
	expectedSum := n * (n + 1) / 2

	for _, num := range nums {
		if seen[num] {
			duplicate = num
		}
		seen[num] = true
		sum += num
	}
	missing = expectedSum - (sum - duplicate)
	return []int{duplicate, missing}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 5, 7, 8}
	exp_sum := 8 * 9 / 2
	var obs_sum int
	for _, v := range arr {
		obs_sum += v
	}
	fmt.Println(exp_sum - obs_sum)
}
