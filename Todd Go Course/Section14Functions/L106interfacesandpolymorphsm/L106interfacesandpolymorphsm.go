package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

type human interface {
	speak()
}

type hotdog int

func (s secretAgent) speak() {
	fmt.Println("I am type secretagent in func / method speak", s.first, s.last, s.ltk)
}

func (s person) speak() {
	fmt.Println("I am type person in func / method speak", s.first, s.last)
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I am func bar using assertion for type person when found", h.(person).first)
	case secretAgent:
		fmt.Println("I am func bar using aseertion for type secretagent when found", h.(secretAgent).first, h.(secretAgent).last, h.(secretAgent).ltk)

	}
	fmt.Println("I am type human", h)
}

func main() {

	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			first: "Miss",
			last:  "Moneypenny",
		},
		ltk: false,
	}

	p1 := person{
		first: "Dr",
		last:  "No",
	}
	fmt.Println(sa1, p1)
	fmt.Println(sa2)
	fmt.Println(p1)
	sa1.speak()
	sa2.speak()
	p1.speak()
	bar(sa1)
	bar(sa2)
	bar(p1)

	// reminder on conversion
	var x hotdog
	x = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int
	y = 54
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}

// Interface says if you have my "method" then you are also my type.
// As person and Secret Agent types also have the method speak then
// they are also of the type human
