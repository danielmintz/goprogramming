package main

import (
	"fmt"
)

var y string

func main() {

	// DECLARE a varaible to be of a certain type
	// and then asign a value of that TYPE to the variable
	fmt.Println("Printing the value of y", y, "ending")
	fmt.Printf("%T\n", y)

	y = "Bond, James Bond"
	fmt.Println(y, "ending")
	fmt.Printf("%T\n", y)
}
