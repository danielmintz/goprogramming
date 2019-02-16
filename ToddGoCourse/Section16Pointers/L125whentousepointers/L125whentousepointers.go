package main

import (
	"fmt"
)

func main() {

	x := 0
	fmt.Println("x befor", &x)
	fmt.Println("x befor", x)
	foo(&x)
	fmt.Println("x after", &x)
	fmt.Println("x after", x)

}

func foo(y *int) {
	fmt.Println("y befor", y)
	fmt.Println("y befor", *y)
	*y = 42
	fmt.Println("y after", y)
	fmt.Println("y after", *y)

}

// When to use pointers:
// Pointers allow you to share a value stored in some memory location. Use pointers when
// you don’t want to pass around a lot of data
// you want to change the data at a location
// Everything in Go is pass by value. Drop any phrases and concepts you may have from other languages.
// Pass by reference, pass by copy - forget those phrases. “Pass by value.”
// That is the only phrase you need to know and remember. That is the only phrase you should use.
// Pass by value. Everything in Go is pass by value. In Go, what you see is what you get (wysiwyg).
// Look at what is happening. That is what you get.
