package main

import (
	"fmt"
)

func maximumUnits(boxTypes [][]int, truckSize int) (int, int) {
	for i := 0; i < len(boxTypes); i++ {
		for j := i + 1; j < len(boxTypes); j++ {
			if boxTypes[j][1] > boxTypes[i][1] {
				tempI := boxTypes[i][0]
				boxTypes[i][0] = boxTypes[j][0]
				boxTypes[j][0] = tempI
				tempJ := boxTypes[i][1]
				boxTypes[i][1] = boxTypes[j][1]
				boxTypes[j][1] = tempJ
			}
		}
	}
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
	// for i := 0; i < len(arr); i++ {
	// 	for j := i + 1; j < len(arr); j++ {
	// 		if arr[j][0] > arr[i][0] {
	// 			tempI := arr[i][0]
	// 			arr[i][0] = arr[j][0]
	// 			arr[j][0] = tempI
	// 			tempJ := arr[i][1]
	// 			arr[i][1] = arr[j][1]
	// 			arr[j][1] = tempJ
	// 		}
	// 	}
	// }
	//fmt.Println(arr[1][1])
	fmt.Println(maximumUnits(arr, 10))
}
