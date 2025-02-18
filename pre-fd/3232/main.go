package main

func canAliceWin(nums []int) bool {
	sdSum := 0
	ddSum := 0
	totalSum := 0
	for _, num := range nums {
		if num < 10 {
			sdSum += num
			totalSum += num
		} else {
			ddSum += num
			totalSum += num
		}
	}

	if totalSum-sdSum < sdSum || totalSum-ddSum < ddSum {
		return true
	} else {
		return false
	}
}
