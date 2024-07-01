package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	// Initialize three pointers:
	// p1 for the last element in the m part of nums1
	// p2 for the last element in nums2
	// p for the last position in nums1 (m + n - 1)
	p1, p2, p := m-1, n-1, m+n-1

	// Iterate while there are elements in both nums1 and nums2
	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}

	// If there are remaining elements in nums2, copy them over
	for p2 >= 0 {
		nums1[p] = nums2[p2]
		p2--
		p--
	}

}
