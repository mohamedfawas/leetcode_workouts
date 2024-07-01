package main

import "fmt"

func countNegatives(grid [][]int) int {
	var neg_count int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] < 0 {
				neg_count++
			}
		}
	}
	return neg_count
}

func main() {
	grid := [][]int{{1, 2, 3, 4}, {4, 5, 6, 7}, {1, 2, 8, 9}}
	fmt.Println(len(grid[0]))
}
