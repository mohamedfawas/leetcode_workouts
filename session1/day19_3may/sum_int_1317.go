package main

import "fmt"

func getNoZeroIntegers(n int) []int {
	var a, b, flag int
	for i := 1; i <= n/2; i++ {
		a, b, flag = i, n-i, 0
		for a > 0 {
			if a%10 == 0 {
				flag = 1
				break
			}
			a /= 10
		}
		if flag == 0 {
			for b > 0 {
				if b%10 == 0 {
					flag = 1
					break
				}
				b /= 10
			}
		}
		if flag == 0 {
			return []int{i, n - i}
		}
	}
	return []int{a, b}
}

func main() {

	fmt.Println(getNoZeroIntegers(9))
}
