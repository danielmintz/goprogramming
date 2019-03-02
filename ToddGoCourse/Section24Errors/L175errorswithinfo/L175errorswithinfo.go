package main

import (
	"errors"
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
		return 0, errors.New("norgate math: square root of negative number")
	}
	return 42, nil
}

// Errors with info
// We can add information to our errors. We can do this with
// errors.New()
// fmt.Errorf()
// builtin.error
// “Error values in Go aren’t special, they are just values like any other, and so you have
// the entire language at your disposal.” - Rob Pike
