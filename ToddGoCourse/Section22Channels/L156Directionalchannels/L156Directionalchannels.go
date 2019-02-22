package main

import "fmt"

func main() {

	ca := make(chan int, 2)
	ca <- 42
	fmt.Println(<-ca)

	ca <- 43
	ca <- 44
	fmt.Println(<-ca)
	fmt.Println(<-ca)

	fmt.Println("----------")
	fmt.Printf("%T\n\n", ca)

	c := make(chan int)
	cr := make(<-chan int) //receive from it
	cs := make(chan<- int) // send onto it

	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", cr)
	fmt.Printf("%T\n", cs)
	//A channel may be constrained only to send or only to receive [general to specific] by conversion
	// or assignment.”
	//"doesn’t assign"
	// Specific to general (doesnt work)
	// c = cr
	// c = cs

	// Specific to Specific doesn't assign (doesn't work)
	// cs = cr

	// general to specific assigns (works)

	//cr = c
	//cs = c

	conversion
	general to specific works
	fmt.Println("-----")
	fmt.Printf("c\t%T\n", (<-chan int)(c))
	fmt.Printf("c\t%T\n", (chan<- int)(c))

	Specific to General doesnt work
	fmt.Println("-----")
	fmt.Printf("c\t%T\n", (chan int)(cr))
	fmt.Printf("c\t%T\n", (chan int)(cs))
}

// Directional channels
// Channel type. Read from left to right.
// “send and receive” means “send and receive”
// send means send
// receive means receive
