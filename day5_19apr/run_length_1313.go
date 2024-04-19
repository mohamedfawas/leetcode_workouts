package main

import "fmt"

func decompressRLElist(nums []int) []int {
	var freq, val int
	var dec_slice []int
	var temp_slice []int
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			freq = nums[i]
			val = nums[i+1]
			temp_slice = make([]int, freq)
			for i := 0; i < freq; i++ {
				temp_slice[i] = val
			}
			dec_slice = append(dec_slice, temp_slice...)
		}
	}
	return dec_slice
}

func main() {
	test_slice := []int{1, 2, 3, 4}
	fmt.Println(decompressRLElist(test_slice))
}
