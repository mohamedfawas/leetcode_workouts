package main

import "fmt"

func findTheDifference(s string, t string) byte {
	var sumS, sumT byte
	for i := 0; i < len(s); i++ {
		sumS += s[i]
	}
	for i := 0; i < len(t); i++ {
		sumT += t[i]
	}
	//fmt.Println(sumT-sumS)
	return sumT - sumS
}

func main() {
	s := "abcd"
	t := "abcde"
	acd := findTheDifference(s, t)
	fmt.Println(acd)

}
