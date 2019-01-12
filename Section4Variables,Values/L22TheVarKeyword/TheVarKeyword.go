package main

import "fmt"

// Declare the variable "y"
// Assign the value 43
// declare and assign = initialization
// var does same as short declaration operator but works outside of a funtion
// body level i.e. it is at package level)

var y = 43

// DECLARE there is a variable with the idenitfier "z"
// and that the variable with the idnintfier "z" is of the type int
// Asigns the ZERO value of the type "int" to  "z"
// false for booleans, 0 for integers, 0.0 for floats, "" for strings
// and nil for pointers, functions, interfaces, slices, channels,  and maps

var z int

// Best Practise limit scope of "var" use and as much as possible use short declarition
// := within the function body

func main() {
	// short declaration variable
	// Declare a Variable and ASIGN a VALUE (of a certin type)
	// short declaration operator does same as var but only within
	// the function body and  from when mentioned.

	x := 42
	fmt.Println(x)

	fmt.Println(y)
	foo()

}

func foo() {
	fmt.Println("again:", y)
	fmt.Println(z)
}
