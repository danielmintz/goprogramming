package main

import (
	"fmt"
)

func main() {
	x := [5]int{5, 4, 3, 2, 1}

	fmt.Println(x)

	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}
