package main

import (
	"fmt"
)

func main() {
	favSport := "MTB"
	switch favSport {
	case "road cycling":
		fmt.Println("nope not mtb")
	case "MTB":
		fmt.Println("yup this is", favSport)
	}

}

// Create a program that uses a switch statement with the switch expression
// specified as a variable of TYPE string with the IDENTIFIER “favSport”.
