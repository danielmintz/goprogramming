package main

import (
	"fmt"
)

func main() {

	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	x := foo(xi...)
	fmt.Println(x)

	xi2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	x2 := bar(xi2)
	fmt.Println(x2)

}

func foo(s ...int) int {
	total := 1
	for _, v := range s {
		total += v
	}
	return total

}

func bar(s []int) int {
	total := 1
	for _, v := range s {
		total += v
	}
	return total

}

//create a func with the identifier foo that
//takes in a variadic parameter of type int
//pass in a value of type []int into your func (unfurl the []int)
//returns the sum of all values of type int passed in

//create a func with the identifier bar that
//takes in a parameter of type []int
//returns the sum of all values of type int passed in
