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
