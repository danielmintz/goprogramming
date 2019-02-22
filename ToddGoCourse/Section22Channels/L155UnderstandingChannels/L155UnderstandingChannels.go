package main

import "fmt"

func main() {

	//doesnt work as blocked (NB channels block, so only work of happening like a relay race - synchronised)
	//-----------------------------------------------------------------------------------------------------
	//ca := make(chan int)
	//ca <- 42
	//fmt.Println(<-ca)

	// Buffered Channels
	// -----------------
	ca := make(chan int, 2) // (buffer allows in exceptions to the lock)
	ca <- 56
	ca <- 102
	fmt.Println(<-ca, <-ca)

	// Using a go routine as a func to work as relay racers passing batton
	// ----------------------------------------------------------------------
	xy := make(chan int)
	go func() {
		xy <- 77

	}()
	fmt.Println(<-xy)

}

//Channels
// Understanding channels
// Channels Introduction
// making a channel
// c := make(chan int)
// putting values on a channel (send onto it)
//	c <- 42
// taking values off of a channel (receive from it)
//		<-c
// buffered channels
//		c := make(chan int, 4)
// channels  block
// they are like runners in a relay race
// they are synchronized
// they have to pass/receive the value at the same time
// just like runners in a relay race have to pass / receive the baton to each other at the same time
// one runner can’t pass the baton at one moment
// and then, later, have the other runner receive the baton
// the baton is passed/received by the runners at the same time
// the value is passed/received synchronously; at the same time
// channels allow us to pass values between goroutines

// “The capacity, in number of elements, sets the size of the buffer in the channel.
// If the capacity is zero or absent, the channel is unbuffered and communication succeeds only
// when both a sender and receiver are ready.” Golang Spec
