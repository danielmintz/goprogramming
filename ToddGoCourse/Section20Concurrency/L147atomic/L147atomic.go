package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Num Go Routine:", runtime.NumGoroutine())

	var counter int64
	const gs = 100
	var wg sync.WaitGroup

	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1) // write to my counter
			runtime.Gosched()
			fmt.Println("Counter:", atomic.LoadInt64(&counter)) //read from my counter

			wg.Done()
		}()
		//fmt.Println("Num Go routine:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("CPU's:", runtime.NumCPU())
	fmt.Println("Num Go routines:", runtime.NumGoroutine())
	fmt.Println("counter:", counter)
}

//Atomic
//We can use package atomic to also prevent a race condition in our incrementer code.
