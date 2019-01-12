package main

import (
	"fmt"
)

func main() {
	i := 1
	for {
		if i > 9 {
			break
		}
		fmt.Println(i)
		i++
	}

}

//There are three ways you can do loops in Go - they all just use the “for” keyword:
//for init; condition; post { }
//for condition { }
//for { }
//Reading documentation for the “for” statement
//language spec
//effective go

//this is for{} and uses neat feature "break"
