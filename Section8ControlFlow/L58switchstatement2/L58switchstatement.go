package main

import (
	"fmt"
)

func main() {
	n := "Bond"
	switch n {
	case "Bond", "Moneypenny", "Dr Q":
		fmt.Println("This is M, J and Q")
	case "Jack":
		fmt.Println("This is just jack")

	}

}

// NB can use case and also variables

// switch statement allows me to switch on to one of these cases evaluating to true
//Conditional - switch statement
//Switch Statements
//switch / case / default statements
//no default fall-through
//creating fall-through
//multiple cases
//cases can be expressions
//evaluate to true, they run

//fallthrough is wonky as prints out when false even though switch only runs when there is a switch.
