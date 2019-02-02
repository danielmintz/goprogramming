package main

import (
	"fmt"
)

func main() {

	c := 42
	fmt.Println(&c)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", &c)

	d := &c
	*d = 64
	fmt.Println(c)

}

// Hands-on exercise #1
// Create a value and assign it to a variable.
// Print the address of that value.
