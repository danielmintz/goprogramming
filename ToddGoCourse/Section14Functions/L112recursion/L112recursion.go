package main

import (
	"fmt"
)

func main() {
	j := factorial(4)
	fmt.Println(j)

	n2 := loop(4)
	fmt.Println(n2)

}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)

}

// doing the same as above using a loop (recommended over a recusrion)
func loop(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}

// Recursion: a func that calls itself

// confused by total
