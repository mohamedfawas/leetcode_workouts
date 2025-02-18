package main

// func findMaxK(nums []int) int {
// 	maxK := -1
// 	recordNums := make(map[int]bool)
// 	for _, num := range nums {
// 		recordNums[num] = true
// 	}

// 	for num, _ := range recordNums {
// 		if recordNums[-num] {
// 			if num > maxK {
// 				maxK = num
// 			}
// 		}
// 	}

// 	return maxK
// }

func findMaxK(nums []int) int {
	maxK := -1
	numRecord := make(map[int]bool)
	for _, num := range nums {
		if numRecord[-num] {
			if num > maxK {
				maxK = num
			}
		} else {
			numRecord[num] = true
		}
	}

	return maxK
}
