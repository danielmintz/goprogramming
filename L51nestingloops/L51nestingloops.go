package main

import (
	"fmt"
)

func main() {
	// for init; conditional; post
	for a := 0; a <= 10; a++ {
		fmt.Printf("the outer loop: %d\n", a)
		for b := 0; b <= 3; b++ {
			fmt.Printf("\t\t the inner loop: %d\n", b)
		}
	}
}
