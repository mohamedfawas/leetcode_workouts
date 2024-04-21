package main

import "fmt"

func judgeCircle(moves string) bool {
	var sumx, sumy int //sum along both the axis
	for _, value := range moves {
		switch value {
		case 'U':
			sumy++
		case 'D':
			sumy--
		case 'L':
			sumx--
		case 'R':
			sumx++
		}
	}
	if sumx == 0 && sumy == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	moves := "LL"
	fmt.Println(judgeCircle(moves))

}
