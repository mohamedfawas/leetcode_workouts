package main

import (
	"slices"
)

func kidsWithCandies(candies []int, extraCandies int) []bool {
	sort_candies := make([]int, len(candies)) // to avoid pass by value, create a new slice and copy values to it
	copy(sort_candies, candies)
	slices.Sort(sort_candies) // sort the array
	max_candies := sort_candies[len(sort_candies)-1]
	for i := 0; i < len(candies); i++ {
		candies[i] = candies[i] + extraCandies
	}
	//fmt.Println("sort candies : ", sort_candies)
	//fmt.Println("candies original after adding extra candies :", candies)
	var bool_arr []bool
	for i := 0; i < len(candies); i++ {
		if candies[i] >= max_candies {
			bool_arr = append(bool_arr, true)
		} else {
			bool_arr = append(bool_arr, false)
		}
	}
	//fmt.Println(bool_arr)
	return bool_arr
}

func main() {
	arr := []int{2, 3, 5, 1, 3}
	// sorted_arr := make([]int, len(arr))
	// copy(sorted_arr, arr)
	// fmt.Println("sorted array before sorting : ", sorted_arr)
	// slices.Sort(sorted_arr)
	// fmt.Println("sorted array after sorting : ", sorted_arr)
	// fmt.Println("original  array after sorting : ", arr)
	kidsWithCandies(arr, 3)
}
