package main

import "fmt"

func checkPerfectNumber(num int) bool {
	var sum int
	for i := 1; i*i < num; i++ {
		if num%i == 0 {
			sum += i
			//fmt.Println("divisor : ", i)
			if i == 1 {
				continue
			} else {
				if (num / i) != i {
					sum += num / i
					//fmt.Println("divisor : ", num/i)
				}
			}
		}
	}
	//fmt.Println("value of sum is : ", sum)
	if sum == num {
		return true
	} else {
		return false
	}
}

func main() {
	num := 28
	fmt.Println(checkPerfectNumber(num))
}
