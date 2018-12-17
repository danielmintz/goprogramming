package main

import (
	"fmt"
)

var y string
var z int

func main() {

	// DECLARE a varaible to be of a certain type
	// and then asign a value of that TYPE to the variable
	fmt.Println("Printing the value of y", y, "ending")
	fmt.Printf("%T\n", y)

	y = "Bond, James Bond"
	fmt.Println(y, "ending")
	fmt.Printf("%T\n", y)

	fmt.Println(z)
	fmt.Printf("%T\n", z)
}

// Use short declaration operator as much as possible
// Use Var at package level and zero value
