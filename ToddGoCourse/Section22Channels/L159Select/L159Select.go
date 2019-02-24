package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	go send(even, odd, quit)

	receive(even, odd, quit)

	fmt.Println("must be done, time to exit")

}

// Send Channel
func send(even, odd chan<- int, quit chan<- bool) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	//close(even)
	//close(odd)
	quit <- true
}

//Receive channel
func receive(even, odd <-chan int, quit <-chan bool) {
	for {
		select {
		case v := <-even:
			fmt.Println("the value received from even is:", v)
		case v := <-odd:
			fmt.Println("the value received from odd is:", v)
		case <-quit:
			return
		}
	}
}

//Select
// Select statements pull the value from whatever channel has a value ready to be pulled.
