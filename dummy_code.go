// this file is for revision practice of the codes I have covered till now
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	names := []string{"A", "B", "C"}
	age := []int{11, 12, 13}

	var people []Person

	for i := 0; i < len(names); i++ {
		person := Person{Name: names[i], Age: age[i]}
		people = append(people, person)
	}

	fmt.Println(people)

}
