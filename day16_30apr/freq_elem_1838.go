package main // incomplete code, do this problem later

import "fmt"

func maxFrequency(nums []int, k int) int {
	frq_map := make(map[int]int)
	for _, v := range nums {
		frq_map[v]++
	}

	//iterate the map, get max freq
	var max_freq, elem_maxf int
	for elem, v := range frq_map {
		if v > max_freq {
			elem_maxf = elem
			max_freq = v
		}
	}
	println(max_freq, elem_maxf)
	return 1
}

func main() {
	slc := []int{1, 2, 4}
	fmt.Println(maxFrequency(slc, 5))
}
