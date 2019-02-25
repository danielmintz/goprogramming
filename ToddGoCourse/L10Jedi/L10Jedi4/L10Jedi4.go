package main

import (
	"fmt"
	"time"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)
	go func() {

		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			c <- i
		}
		q <- 1
		close(c)

	}()
	return c
}

func receive(c, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println("value of c is:", v)
		case <-q:
			return

		}
	}

}

// Hands-on exercise #4
//Starting with this code, pull the values off the channel using a select statement
