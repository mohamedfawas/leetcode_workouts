package main

import "fmt"

func returnToBoundaryCount(nums []int) int {
	var sum, count int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if i > 0 {
			if sum == 0 {
				count++
			}
		}
	}
	return count
}

func main() {
	arr := []int{2, 3, -5}
	fmt.Println("test array is : ", arr)
	fmt.Println("no: of times ant returns to boundary is : ", returnToBoundaryCount(arr))
}
