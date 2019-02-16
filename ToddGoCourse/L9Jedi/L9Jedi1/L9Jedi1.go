package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("CPU's:", runtime.NumCPU())
	fmt.Println("num go routines:", runtime.NumGoroutine())
	wg.Add(3)
	bar()
	go foo()
	go foobar("It's me")
	go func() {
		fmt.Println("we are number 4")
		wg.Done()
	}()
	go func() {
		fmt.Println("we are number 5")
		wg.Done()
	}()

	go func() {
		fmt.Println("we are number 6")
		wg.Done()
	}()
	go func() {
		fmt.Println("we are number 7")
		wg.Done()
	}()
	go func() {
		fmt.Println("we are number 8")
		wg.Done()
	}()
	go func() {
		fmt.Println("we are number 9")
		wg.Done()
	}()

	fmt.Println("num go routines:", runtime.NumGoroutine())
	wg.Wait()

}

func bar() {
	for x := 0; x < 10; x++ {
		fmt.Println("bar:", x)
	}

}

func foo() {
	for y := 0; y < 10; y++ {
		fmt.Println("foo:", y)
	}
	wg.Done()

}

func foobar(c string) {
	fmt.Println(c)
	wg.Done()
}

//Hands-on exercise #1
//in addition to the main goroutine, launch two additional goroutines
//each additional goroutine should print something out
//use waitgroups to make sure each goroutine finishes before your program exists
