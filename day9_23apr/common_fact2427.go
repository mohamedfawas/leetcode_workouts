package main

import "fmt"

func commonFactors(a int, b int) int {
	var count int
	for i := 1; i <= min(a, b); i++ {
		if a%i == 0 && b%i == 0 {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println("test")
}
