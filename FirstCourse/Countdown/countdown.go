package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 99; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(time.Second)

	}
	fmt.Println("Happy New Year")

}
