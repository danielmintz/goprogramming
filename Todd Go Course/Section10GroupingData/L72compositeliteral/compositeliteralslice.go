package main

import (
	"fmt"
)

func main() {
	//COMPOSITE LITERAL
	x := []int{5, 4, 6, 7, 89}
	fmt.Println(x)

}

//A SLICE allows you to group VALUES together of the same TYPE
//Composite literals construct values for structs, arrays, slices, and maps and create a new value each time
//they are evaluated.  They consist of the type of the literal followed by a brace-bound list of elements

//A SLICE holds VALUES of the same TYPE. If I wanted to store all of my favorite numbers,
//I would use a slice of int. If I wanted to store all of my favorite foods, I would use a slice of string.
//We will use a COMPOSITE LITERAL to create a slice. A composite literal is created by having the TYPE
//followed by CURLY BRACES and then putting the appropriate values in the curly brace area.
