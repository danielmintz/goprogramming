package main

import (
	"fmt"
	"log"
)

func main() {

	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("cant do a sqaure root of a negative such as: %v\t%T", f, f)
	}
	return 42, nil
}
