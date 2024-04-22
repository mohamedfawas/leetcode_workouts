package main

import "fmt"

func isPerfectSquare(num int) bool {
	for i := 1; i <= num; i++ {
		if i*i == num {
			return true
		} else if i*i > num {
			return false
		}
	}
	return false
}

func main() {

	fmt.Println("Give an input integer to test: ")
	var num int
	fmt.Scan(&num)
	fmt.Println(isPerfectSquare(num))

}
