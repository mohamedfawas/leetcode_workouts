package main

import (
	"fmt"
)

func numIdenticalPairs(nums []int) int {
	var count int
	numsMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		numsMap[nums[i]]++
	}
	fmt.Println(numsMap)
	for _, element := range numsMap {
		fmt.Println("element ; ", element)
		count += element * (element - 1) / 2
	}
	return count
}

func main() {
	nums1 := []int{1, 2, 3, 1, 1, 3, 4}
	// numIdenticalPairs(nums1)
	fmt.Println(numIdenticalPairs(nums1))
}
