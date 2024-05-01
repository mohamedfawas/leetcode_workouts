package main

import (
	"fmt"
	"strconv"
)

func alternateDigitSum(n int) int {
	var l_digit, sum, sign int
	var temp int = n
	sign = 1
	if len(strconv.Itoa(temp))%2 == 0 { // when lenght of the number is even
		sign = -1
	}
	for temp > 0 {
		l_digit = temp % 10
		temp /= 10
		sum += l_digit * sign
		sign *= -1
	}
	return sum
}

func main() {
	fmt.Println(alternateDigitSum(25))
}
