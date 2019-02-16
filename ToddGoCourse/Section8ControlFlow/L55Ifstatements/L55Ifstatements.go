package main

import (
	"fmt"
)

func main() {

	if true {
		fmt.Println("001")
	}

	if false {
		fmt.Println("002")
	}

	if !true {
		fmt.Println("003")
	}

	if !false {
		fmt.Println("004")
	}

	if 45 == 56 {
		fmt.Println("005")
	}

	if 45 <= 56 {
		fmt.Println("006")
	}

	if 45 != 56 {
		fmt.Println("007")
	}

	if !(2 == 2) {
		fmt.Println("008")
	}

	if !(2 != 2) {
		fmt.Println("009")
	}

}

//3 types of statements: sequence, iterative (loop) & conditional. This one is all about conditional
// Conditional if statement:
//If Statements
//bool
//true
//false
//the not operator
//!
//initialization statement
