package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			time.Sleep(time.Millisecond * 200)
			c <- i
		}
		close(c)

	}()

	for v := range c {
		fmt.Println(v)
	}
	fmt.Println("yo we are there,")

}

// Range
// Range stops reading from a channel when the channel is closed.
