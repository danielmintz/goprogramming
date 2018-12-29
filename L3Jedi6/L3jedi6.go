package main

import (
	"fmt"
)

func main() {
	x := "James oooBond"
	if x == "Miss Moneypenny" {
		fmt.Println("not this one")
	} else if x == "Mr Q" {
		fmt.Println("could be this")
	} else {
		fmt.Println("it's not bond, nope")
	}
}

//Create a program that shows the “if statement” in action.
//Building on the previous hands-on exercise, create a program that uses “else if” and “else”.
