package main

import "fmt"

func main() {

	x := foo()
	fmt.Println(x())

	y := bar("amazeballs")
	y()

}

func foo() func() int {
	return func() int {
		return 451
	}
}

func bar(s string) func() string {
	return func() string {
		fmt.Println("yo func that returns a funk", s)
		return s
	}
}

//Hands-on exercise #8
//Create a func which returns a func
//assign the returned func to a variable
//call the returned func
