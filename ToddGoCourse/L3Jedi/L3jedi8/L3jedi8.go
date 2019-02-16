package main

import (
	"fmt"
)

func main() {
	favSport := "surfing"
	switch favSport {
	case "icehockey":
		fmt.Println("not quite sport but close")
	case "surfing":
		fmt.Println("you got it")
	}

}

//Create a program that uses a switch statement with the switch expression specified
// as a variable of TYPE string with the IDENTIFIER “favSport”.
