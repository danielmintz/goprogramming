package main

import (
	"fmt"
)

func main() {

	a := foo()
	b := foo()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())

}

func foo() func() int {
	x := 0
	return func() int {

		x++
		return x

	}

}

//Closure is when we have “enclosed” the scope of a variable in some code block.
//For this hands-on exercise, create a func which “encloses” the scope of a variable:
