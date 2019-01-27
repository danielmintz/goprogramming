package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.length
}

func (s circle) area() float64 {
	return math.Pi * s.radius * s.radius
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println("this is the area:", s.area())

}

func main() {

	squ := square{
		length: 12.2,
	}

	cir := circle{
		radius: 8.1,
	}

	info(squ)
	info(cir)

}

// create a type SQUARE
// create a type CIRCLE
// attach a method to each that calculates AREA and returns it
// circle area= π r 2
// square area = L * W
// create a type SHAPE that defines an interface as anything that has the AREA method
// create a func INFO which takes type shape and then prints the area
// create a value of type square
// create a value of type circle
// use func info to print the area of square
// use func info to print the area of circle
