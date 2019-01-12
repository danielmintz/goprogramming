package main

import (
	"fmt"
)

func main() {
	n := "Bond"
	switch n {
	//switch "Bond" {
	case "Moneypenny":
		fmt.Println("this is Moneypenny")
	case "Bond":
		fmt.Println("this is James, James Bond")
	}
}

// NB can use case and also variables
