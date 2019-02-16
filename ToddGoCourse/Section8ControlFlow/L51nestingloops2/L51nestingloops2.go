package main

import (
	"fmt"
)

func main() {
	for a := 0; a <= 10; a++ {
		for b := 0; b < 3; b++ {
			fmt.Printf("The outer loop: %d\t The inner loop: %d\n", a, b)
		}
	}
}
