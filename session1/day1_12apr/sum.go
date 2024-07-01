package main

import "fmt"

func main() {
	fmt.Print("Give an input number num1 :")
	var num1 int
	fmt.Scan(&num1)
	fmt.Print("Give an input number num2 :")
	var num2 int
	fmt.Scan(&num2)
	var sum int
	sum = num1 + num2
	fmt.Println("Sum of the given two numbers are : ", sum)
}
