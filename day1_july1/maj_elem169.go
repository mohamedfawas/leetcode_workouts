package main

func majorityElement(nums []int) int {
	var el, count int
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			count = 1
			el = nums[i]
		} else if nums[i] == el {
			count++
		} else {
			count--
		}
	}
	return el
}
