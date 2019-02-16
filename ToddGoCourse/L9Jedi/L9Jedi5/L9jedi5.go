package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	var incrementor int64
	const gs = 100
	var wg sync.WaitGroup

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {

			atomic.AddInt64(&incrementor, 1) //add to counter
			runtime.Gosched()
			fmt.Println("counter", atomic.LoadInt64(&incrementor)) //reads from counter
			wg.Done()

		}()

	}
	wg.Wait()
	fmt.Println("count", incrementor)

}

//Hands-on exercise #5
//Fix the race condition you created in exercise #4 by using package atomic
