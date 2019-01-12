package main

import (
	"fmt"
)

func main() {

	p1 := struct {
		first   string
		last    string
		age     int
		hobbies []string
		number  map[string]int
	}{
		first: "James",
		last:  "Bond",
		age:   32,
		hobbies: []string{
			"MTB",
			"4X",
			"Skiing",
		},
		number: map[string]int{
			"JB": 333,
			"MP": 555,
		},
	}
	fmt.Println(p1.first)
	fmt.Println(p1.last)
	fmt.Println(p1.age)
	for i, v := range p1.hobbies {
		fmt.Println(i, v)
	}
	for i, v := range p1.number {
		fmt.Println(i, v)
	}
}
