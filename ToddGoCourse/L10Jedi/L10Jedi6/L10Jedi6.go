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
}

// Hands-on exercise #6
// write a program that puts 100 numbers to a channel
// pull the numbers off the channel and print them
