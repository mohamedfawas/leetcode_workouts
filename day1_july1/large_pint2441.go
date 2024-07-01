package main

import "fmt"

//In Go, the struct{} type is an empty struct that takes up zero bytes of storage.
//It's commonly used in maps to represent sets, where only the keys are important and the values are irrelevant.

func findMaxK(nums []int) int {
	numSet := make(map[int]struct{})
	maxK := -1
	for _, num := range nums {
		numSet[num] = struct{}{}
	}

	for _, num := range nums {
		if num > 0 {
			if _, found := numSet[-num]; found {
				maxK = max(maxK, num)
			}
		}
	}
	return maxK
}

func main() {
	arr := []int{1, -1, 3, -3, 4, 5, 6, -8, 9}
	numFreq := make(map[int]struct{})
	maxK := -1
	for _, v := range arr {
		numFreq[v] = struct{}{}
	}
	fmt.Println(numFreq)
	for _, v := range arr {
		fmt.Println(numFreq[-v])
	}

	for _, num := range arr {
		if num > 0 {
			fmt.Println(numFreq[-num])
		}
	}
	fmt.Println(maxK)
}
