package leetcode

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) (ans []int) {
	count := make(map[int]int)

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

	result := []int{}
	for num, cnt := range count {
		if cnt >= 2 {
			result = append(result, num)
		}
	}
	return result
}
