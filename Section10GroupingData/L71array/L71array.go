package main

import (
	"fmt"
)

var x [5]int

func main() {
	fmt.Println(x)
	x[3] = 42
	fmt.Println(x)
	fmt.Println(len(x))

}

// zero based index
//Arrays are mostly used as a building block in the Go programming language.
//In some instances, we might use an array, but mostly the recommendation is to use slices instead.
// array
// a numbered sequence of elements of a single type
// does not change in size
// used for Go internals; generally not recommended for your code
// https://golang.org/ref/spec#Array_types
