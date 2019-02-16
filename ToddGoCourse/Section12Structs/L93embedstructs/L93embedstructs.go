package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	licensetokill bool
}

func main() {

	p2 := person{
		first: "Debra",
		last:  "Scrivner",
		age:   32,
	}
	p3 := person{
		first: "Marni",
		last:  "Humphries",
		age:   13,
	}
	sa1 := secretAgent{
		person: person{
			first: "Daniel",
			last:  "Mintz",
			age:   48,
		},
		licensetokill: true,
	}
	x := "one day met"
	y := "and thinks he is a secret agent with a license to kill,"
	z := "and she had a lovely daughter called"
	a := "was a bit of a misery"
	b := "and he fell in love with"

	fmt.Println(sa1.first, sa1.last, x, p2.first, p2.last, z, p3.first, p3.last, b, p2.first, ".", sa1.first, a, y)
}

//embedding one type in another type the inner gets promoted ot the outer type
