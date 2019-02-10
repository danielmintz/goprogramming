package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println("area:", s.area())
}

func main() {

	c := circle{
		radius: 5,
	}
	fmt.Println(c.area())
	info(&c)
	//info(c) // wont work as methos set is a pointer

}

//“The method set of a type determines the INTERFACES that the type implements.....” ~
