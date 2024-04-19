package main

import "fmt"

func sumOfTheDigitsOfHarshadNumber(x int) int {
	var count, sum int
	for i := 1; i*i <= x; i++ {
		if x%i == 0 {
			count++
			sum += i
			if (x / i) != i {
				count++
				sum += int(x / i)
			}
		}
	}
	if count > 0 {
		return sum
	} else {
		return -1
	}
}

func main() {
	fmt.Println("test")
}
