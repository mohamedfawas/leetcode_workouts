package main

import "fmt"

func isPowerOfFour(n int) bool {
	if n == 1 {
		return true
	} else if n%4 != 0 || n <= 0 {
		return false
	}
	return isPowerOfFour(n / 4)
}

func main() {
	fmt.Println()
}
