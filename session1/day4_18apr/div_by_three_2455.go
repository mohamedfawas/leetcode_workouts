package main

import "fmt"

func averageValue(nums []int) int {
	var sum, count int
	for i := 1; i <= len(nums); i++ {
		if nums[i-1]%3 == 0 && nums[i-1]%2 == 0 {
			sum += nums[i-1]
			count++
		}
	}
	fmt.Println("sum is : ", sum)
	fmt.Println("count is : ", count)
	//avg_val = sum/count
	if count == 0 {
		return 0
	} else {
		return (sum / count)
	}
}

func main() {
	arr := []int{1, 3, 6, 10, 12, 15}
	fmt.Println(arr)
	fmt.Println(averageValue(arr))
	arr2 := []int{4, 4, 9, 10}
	fmt.Println(arr2)
	fmt.Println(averageValue(arr2))
}
