package main

import (
	"fmt"
)

func main() {

	f := func() {
		fmt.Println("my first func expression")
	}
	f()

	g := func(x int) {
		fmt.Println("my seodn fucn expression with more stuff like george orwell year:", x)
	}
	g(1984)
}
