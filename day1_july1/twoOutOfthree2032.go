package main

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) (ans []int) {
	// Create a map to count occurrences in different arrays
	//We create a map count to keep track of in how many arrays each number appears.
	count := make(map[int]int)

	// Helper function to update the count
	//We define a helper function updateCount that takes an array and updates the count map.
	//It uses a local seen map to ensure that each number is counted only once per array.
	updateCount := func(nums []int) {
		seen := make(map[int]bool)
		for _, num := range nums {
			if !seen[num] {
				count[num]++
				seen[num] = true
			}
		}
	}

	updateCount(nums1)
	updateCount(nums2)
	updateCount(nums3)

	// Collect numbers that appear in at least two arrays
	//Finally, we iterate through the count map and add any number that appears in at least two arrays
	//(i.e., has a count of 2 or more) to the result slice.
	result := []int{}
	for num, c := range count {
		if c >= 2 {
			result = append(result, num)
		}
	}
	return result
}
