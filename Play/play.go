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
	fmt.Println(s.first, s.last, s.ltk)
}

func (s person) speak() {
	fmt.Println(s.first, s.last)
}

func bar(r human) {
	switch r.(type) {
	case person:
		fmt.Println("I am ", r.(person).first, r.(person).last)
	case secretAgent:
		fmt.Println("I am ", r.(secretAgent).first, r.(secretAgent).last, r.(secretAgent).ltk)
	}

}

func main() {

	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "bond",
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

	sa1.speak()
	sa2.speak()
	p1.speak()
	bar(sa2)
	bar(sa1)
	bar(p1)

	var x hotdog
	x = 42
	fmt.Printf("%T\n", x)
	fmt.Println(x)

	y := 53
	fmt.Printf("%T\n", y)
	fmt.Println(y)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	fmt.Println(x)
	fmt.Printf("%T\n", x)

}

//func (r receiver) identifier(parameters) (returns) { code }

//You can create a func which takes an unlimited number of arguments.
// When you do this, this is known as a “variadic parameter.” When use the lexical element operator “...T”
//to signify a variadic parameter (there “T” represents some type).
