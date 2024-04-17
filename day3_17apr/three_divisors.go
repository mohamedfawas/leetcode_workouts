package main

import "fmt"

func isThree(n int) bool {
	var count int
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			count++
			if n/i != i {
				count++
			}
		}
	}

	if count == 3 {
		return true
	} else {
		return false
	}
}

func main() {
	var num int
	fmt.Println("Give the input number to check for no: of divisors : ")
	fmt.Scan(&num)

	result := isThree(num)
	fmt.Println("given number have exactly three divisors : ", result)
}
