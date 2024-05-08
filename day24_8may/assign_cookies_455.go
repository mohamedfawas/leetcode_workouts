package main

import (
	"fmt"
	"slices"
)

func findContentChildren(g []int, s []int) int {
	slices.Sort(g)
	slices.Sort(s)
	var i, j int
	for i < len(g) && j < len(s) {
		if s[j] >= g[i] {
			i++
		}
		j++
	}
	return i
}

func main() {
	arr1 := []int{10, 9, 8, 7}
	arr2 := []int{5, 6, 7, 8}
	fmt.Println(findContentChildren(arr1, arr2))
}
