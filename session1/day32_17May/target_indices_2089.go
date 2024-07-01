package main

import "slices"

func targetIndices(nums []int, target int) []int {
	slices.Sort(nums)
	indexArr := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			indexArr = append(indexArr, i)
		}
	}
	return indexArr
}

func main() {

}
