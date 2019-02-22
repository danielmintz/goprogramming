package main

import "fmt"

func main() {

	c := make(chan int)

	go foo(c)

	bar(c)

	fmt.Printf("this is the exit")
}
func foo(c chan<- int) {
	c <- 42
}

func bar(c <-chan int) {
	fmt.Println("this is the value", <-c)
}

// loops
// init, cond, post
