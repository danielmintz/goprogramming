package main

import (
	"fmt"
)

func main() {
	switch {
	case false:
		fmt.Println("this is false so wont print")
	case true:
		fmt.Println("this is true so will print")
		fallthrough
	case 1+1 == 2:
		fmt.Println("this is true but wont print as already found a true statment unless you use a fallthrough statement")
		fallthrough
	case 1 == 3:
		fmt.Println("this is false but prints if a falltorugh statemnet used above")
		fallthrough
	default:
		fmt.Println("this is the default")
	}
}

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
