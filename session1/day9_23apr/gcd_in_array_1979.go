package main

import (
	"fmt"
	"slices"
)

func findGCD(nums []int) int {
	min, max := slices.Min(nums), slices.Max(nums)
	for i := min; i > 1; i-- {
		if min%i == 0 && max%i == 0 {
			return i
		}
	}
	return 1
}

func main() {
	arr1 := []int{2, 5, 6, 9, 10}
	fmt.Println(findGCD(arr1))
}
