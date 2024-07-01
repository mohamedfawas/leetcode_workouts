package main

import "fmt"

func subtractProductAndSum(n int) int {
	var product, sum, l_digit int
	product = 1
	for n > 0 {
		l_digit = n % 10
		product *= l_digit
		sum += l_digit
		n /= 10
	}
	return (product - sum)
}

func main() {
	fmt.Println()
}
