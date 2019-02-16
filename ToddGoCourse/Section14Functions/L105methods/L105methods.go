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

// func (r receiver) identifier (parameters) (return(s)) { code  }

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, s.ltk)
}

func main() {

	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}
	sa2 := secretAgent{
		person: person{
			"Miss",
			"Moneypenny",
		},
		ltk: false,
	}
	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()
}

//A method is nothing more than a FUNC attached to a TYPE.
//When you attach a func to a type it is a method of that type.
//You attach a func to a type with a RECEIVER.

//We have  attached the function 'speak()' to  any value of the type 'secretAgent' so that any value of that type e.g. 'sa1'
// has access to the method by chaining it with a dot i.e. we chain the value to the function e.g.  sa1.speak()
