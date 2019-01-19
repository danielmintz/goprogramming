package main

import (
	"fmt"
)

func main() {

	x := sum(2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("the total is", x)
}

func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("the value at position", i, "which is", v, "the total is now", sum)
	}
	return sum
}

//You can create a func which takes an unlimited number of arguments.
// When you do this, this is known as a “variadic parameter.” When use the lexical element operator “...T”
//to signify a variadic parameter (there “T” represents some type).
