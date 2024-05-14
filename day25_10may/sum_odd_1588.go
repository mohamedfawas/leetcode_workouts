package main

import "fmt"

func sumOddLengthSubarrays(arr []int) int {
	total_sum := 0
	l := len(arr)
	for i := 0; i < l; i++ {
		for j := i; j < l; j += 2 {
			for k := i; k <= j; k++ {
				total_sum += arr[k]
			}
		}
	}
	return total_sum
}

func main() {
	fmt.Println()
}
