package main

import "fmt"

func main() {

	c := make(chan int)

	go foo(c)

	for v := range c {
		fmt.Println(v)
	}

	fmt.Printf("this is the exit")
}
func foo(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)

}

// loops
// init, cond, post
