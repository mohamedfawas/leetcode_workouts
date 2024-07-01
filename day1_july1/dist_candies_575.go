package main

import "fmt"

func distributeCandies(candyType []int) int {
	uniqueCandies := make(map[int]bool)
	for _, candy := range candyType {
		uniqueCandies[candy] = true
	}

	numUniqueCandies := len(uniqueCandies)
	maxNumEat := len(candyType) / 2
	if numUniqueCandies < maxNumEat {
		return numUniqueCandies
	}
	return maxNumEat
}

func exampleProblem(candyType []int) {
	uniqueCandies := make(map[int]bool)
	for _, v := range candyType {
		uniqueCandies[v] = true
	}
	fmt.Println(uniqueCandies)
}

func main() {
	fmt.Println("example of how map works ")
	candyType := []int{1, 1, 2, 2, 3, 3}
	exampleProblem(candyType)
	fmt.Println("=====================")

	fmt.Println(distributeCandies(candyType))
}
