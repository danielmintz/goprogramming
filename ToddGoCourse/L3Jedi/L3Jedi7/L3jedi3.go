package main

import (
	"fmt"
)

func main() {
	switch {
	case 2 != 2:
		fmt.Println("002")
	case 2 == 2:
		fmt.Println("001")
	}

}

//Create a program that uses a switch statement with no switch expression specified.
