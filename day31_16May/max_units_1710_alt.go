package main

import (
	"fmt"
	"sort"
)

func maximumUnits(boxTypes [][]int, truckSize int) (int, int) {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1] // return in decreasing order
	})
	var sumBox, noOfUnits int
	for i := 0; i < len(boxTypes); i++ {
		if sumBox < truckSize {
			spaceAvailable := truckSize - sumBox
			if spaceAvailable < boxTypes[i][0] {
				sumBox += spaceAvailable
				noOfUnits += spaceAvailable * boxTypes[i][1]
			} else {
				sumBox += boxTypes[i][0]
				noOfUnits += boxTypes[i][0] * boxTypes[i][1]
			}
		} else {
			break
		}
	}
	fmt.Println(boxTypes)
	return noOfUnits, sumBox
}

func main() {
	arr := [][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}}
	fmt.Println(maximumUnits(arr, 10))
}
