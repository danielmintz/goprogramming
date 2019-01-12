package main

import (
	"fmt"
)

type person struct {
	first   string
	last    string
	favicec []string
}

func main() {

	p1 := person{
		first: "Dan",
		last:  "Mintz",
		favicec: []string{
			"strawberry",
			"chocolate",
			"cream",
		},
	}
	p2 := person{
		first: "Debs",
		last:  "Scrivs",
		favicec: []string{
			"vanilla",
			"coffee",
		},
	}

	fmt.Println(p1.first)
	for i, v := range p1.favicec {
		fmt.Println(i, v)
	}
	fmt.Println(p2.first)
	for i, v := range p2.favicec {
		fmt.Println(i, v)
	}
}

//Print out the values, ranging over the elements in the slice which stores the favorite flavors.
