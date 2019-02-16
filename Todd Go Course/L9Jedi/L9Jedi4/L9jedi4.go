package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	incrementor := 0
	const gs = 100
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := incrementor
			//runtime.Gosched()
			v++
			incrementor = v
			mu.Unlock()
			wg.Done()

		}()

		fmt.Println("go routines:", runtime.NumGoroutine())

	}
	wg.Wait()
	fmt.Println("count", incrementor)

}

//Hands-on exercise #4
//Fix the race condition you created in the previous exercise by using a mutex
//it makes sense to remove runtime.Gosched()
