package main

import (
	"fmt"
)

func main() {
	i := 0
	for i < 9 {
		fmt.Println(i)
		i++
	}
	fmt.Println("done")
}

//There are three ways you can do loops in Go - they all just use the “for” keyword:
//for init; condition; post { }
//for condition { }
//for { }
//Reading documentation for the “for” statement
//language spec
//effective go

//this one is the for condition {}
