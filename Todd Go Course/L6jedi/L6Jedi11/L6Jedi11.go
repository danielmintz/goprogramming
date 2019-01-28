package main

import (
	"fmt"
)

func main() {

	a := factorial(5)
	fmt.Println(a)

	b := Loop(6)
	fmt.Println(b)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func Loop(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}

//Recursion / factorial ie a func that call itseld
