package main

import (
	"fmt"
)

func main() {

	const (
		a = 2018 + iota
		b
		c
		d
	)
	fmt.Println(a, b, c, d)
}

//Using iota, create 4 constants for the NEXT 4 years. Print the constant values.
