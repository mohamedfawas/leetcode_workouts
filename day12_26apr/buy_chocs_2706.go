package main

import (
	"fmt"
	"slices"
)

func buyChoco(prices []int, money int) int {
	slices.Sort(prices)
	sum := prices[0] + prices[1]
	if sum > money {
		return money
	} else {
		return money - sum
	}
}

func main() {
	arr := []int{1, 2, 2}
	money := 3
	fmt.Println(buyChoco(arr, money))
}
