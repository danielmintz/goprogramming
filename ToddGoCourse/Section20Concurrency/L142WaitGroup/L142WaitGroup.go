package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPU\t\t", runtime.NumCPU())
	fmt.Println("Go Routines\t", runtime.NumGoroutine())

	wg.Add(1)
	go foo()
	bar()

	fmt.Println("CPU\t\t", runtime.NumCPU())
	fmt.Println("Go Routines\t", runtime.NumGoroutine())
	wg.Wait()

}

func foo() {
	for x := 0; x < 10; x++ {
		fmt.Println("Foo", x)
	}
	wg.Done()
}

func bar() {
	for y := 0; y < 10; y++ {
		fmt.Println("bar", y)
	}

}

//WaitGroup
// A WaitGroup waits for a collection of goroutines to finish.
// The main goroutine calls Add to set the number of goroutines to wait for.
// Then each of the goroutines runs and calls Done when finished.
// At the same time, Wait can be used to block until all goroutines have finished.
// Writing concurrent code is super easy: all we do is put “go” in front of a function or method call.

// runtime.NumCPU()
// runtime.NumGoroutine()
// sync.WaitGroup
// func (wg *WaitGroup) Add(delta int)
// func (wg *WaitGroup) Done()
// func (wg *WaitGroup) Wait()
