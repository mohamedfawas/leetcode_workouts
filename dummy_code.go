package main

import "fmt"

func main() {
	arr1 := []string{"a", "r", "s", "e", "h"}
	arr2 := []int{1, 2, 3, 4, 5}
	mp1 := make(map[string]int)
	for i := 0; i < len(arr1); i++ {
		mp1[arr1[i]] = arr2[i]
	}
	fmt.Println(mp1)
}
