package main

import (
	"fmt"
)

var y = 42

// DECLARE the VARIABLE with the IDENTIFIER type Z
// is of type string
// and I ASSIGN the VALUE of "Shaken not stirred"

var z = "Shaken not stirred"
var a string = `James said "Shaken, 

not Stirred"`

// This is a static programming language
// A Variable is DECLARED to hold a VALUE of a certain TYPE
// Not a DYNAMIC programming language

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y) // Prints out the type
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(a)
	fmt.Printf("%T\n", a)

}
