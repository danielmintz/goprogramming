package main

import (
	"fmt"
)

type car struct {
	Make    string
	Model   string
	RegYear int
	Class   []string
}

func (s car) driveway() {
	fmt.Println("I want to own a", s.Make, s.Model, s.RegYear)
}

func main() {

	p1 := car{
		Make:    "Volkswagen",
		Model:   "Golf",
		RegYear: 2012,
		Class: []string{
			"TopRange",
			"MiddleRange",
			"BottomRange",
		},
	}
	fmt.Println(p1)
	p1.driveway()
}

//number  map[string]int
