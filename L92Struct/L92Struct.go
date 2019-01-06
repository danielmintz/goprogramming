package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func main() {

	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}
	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		age:   24,
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)
}

//A struct is an composite data type. (composite data types, aka, aggregate data types, aka, complex data types).
// Structs allow us to compose together values of different types.
// we have created values (feilds) of a type (struct)
//composing together values of differenet type
