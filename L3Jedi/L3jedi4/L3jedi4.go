package main

import (
	"fmt"
)

func main() {
	bd := 1970
	for {
		if bd > 2018 {
			break
		}
		fmt.Println(bd)
		bd++

	}

}

//Create a for loop using this syntax
//for { }
//Have it print out the years you have been alive.
