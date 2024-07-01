package main

import "fmt"

/*
Given an integer array nums, return true if any value appears at least twice
 in the array, and return false if every element is distinct.
*/

func checkVals(array []int) bool {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] == array[j] {
				return true
			}
		}
	}
	return false
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 2}
	fmt.Println(checkVals(arr))
}
