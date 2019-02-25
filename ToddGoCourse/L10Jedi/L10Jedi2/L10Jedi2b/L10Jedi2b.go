package main

import (
	"fmt"
)

func main() {
	cr := make(chan int)

	go func() {
		cr <- 42
	}()
	fmt.Println(<-cr)

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)
}

//Hands-on exercise #2
// Get this code running:

// https://play.golang.org/p/_DBRueImEq
// solution: https://play.golang.org/p/mgw750EPp4
