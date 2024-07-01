package main

func sumOfTheDigitsOfHarshadNumber(x int) int {
	var initial_int, last_digit, sum int
	initial_int = x
	for x > 0 {
		last_digit = x % 10
		sum += last_digit
		x /= 10
	}
	if initial_int%sum == 0 {
		return sum
	}
	return -1
}
