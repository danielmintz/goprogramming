package main

import (
	"fmt"
)

func main() {

	a := foo()
	b, c := bar()

	fmt.Println(a, b, c)

}

func foo() int {
	return 42
}

func bar() (int, string) {
	return 617, "yooooo"
}

//create a func with the identifier foo that returns an int
//create a func with the identifier bar that returns an int and a string
//call both funcs
//print out their results
