package main

import "math"

func thirdMax(nums []int) int {
	m1, m2, m3 := math.MinInt, math.MinInt, math.MinInt
	for _, num := range nums {
		if num == m1 || num == m2 || num == m3 {
			continue
		}
		if num > m1 {
			m3, m2, m1 = m2, m1, num
		} else if num > m2 {
			m3, m2 = m2, num
		} else if num > m3 {
			m3 = num
		}
	}
	if m3 != math.MinInt {
		return m3
	}
	return m1
}
