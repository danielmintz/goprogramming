package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("CPU's", runtime.NumCPU())
	//fmt.Println("GoRoutines:", runtime.NumGoroutine())

	Counter := 0
	const gs = 100
	var wg sync.WaitGroup

	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {

		go func() {
			mu.Lock()
			v := Counter
			runtime.Gosched()
			v++
			Counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("GoRoutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("GoRoutines:", runtime.NumGoroutine())
	fmt.Println("count:", Counter)

}

//Race condition
//Here is a picture of the race condition we are going to create. Race conditions are not good code. 
//A race condition will give unpredictable results. We will see how to fix this race condition in the next video.
//code: https://play.golang.org/p/FYGoflKQej
//video: 134
//Mutex
//A “mutex” is a mutual exclusion lock. Mutexes allow us to lock our code so that only one 
//goroutine can access that locked chunk of code at a time.
